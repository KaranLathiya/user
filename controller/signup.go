package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

func (c *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	var signup request.Signup
	err := utils.BodyReadAndValidate(r.Body, &signup, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	userExistence, err := c.repo.UserExistence(signup.Email, signup.PhoneNumber)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	if userExistence {
		error_handling.ErrorMessageResponse(w, error_handling.UserAlreadyExist)
		return
	}
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.OTPGenerateError)
		return
	}
	hashedOTP, err := utils.Bcrypt(otp)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.BcryptError)
		return
	}
	storeOTP := request.StoreOTP{
		Email:       signup.Email,
		PhoneNumber: signup.PhoneNumber,
		CountryCode: signup.CountryCode,
		EventType:   "signup",
		LoginType:   signup.LoginType,
		HashedOTP:   hashedOTP,
	}
	err = c.repo.StoreOTP(storeOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP for signup: "
	if signup.LoginType == "email" {
		go utils.SendOTPEmail(signup.Email, otp, subject)
	} else {
		go utils.SendOTPPhone(signup.CountryCode, signup.PhoneNumber, otp, subject)
	}
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.OTP_SENT}
	utils.SuccessMessageResponse(w, 200, successResponse)
}
