package controller

import (
	"user/constant"
	"net/http"
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
	otp, err := utils.GenerateOTP(6)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.OTPGenerateError)
		return
	}
	hashedOTP, err := utils.Bcrypt(otp)
	if err != nil{
		error_handling.ErrorMessageResponse(w, error_handling.BcryptError)
		return
	}
	err = c.repo.StoreOTP(signup,hashedOTP, "signup")
	if err != nil{
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject:="OTP for signup: "
	if signup.LoginType == "email" {
		go utils.SendOTPEmail(signup.Email,otp,subject)
	} else {
		go utils.SendOTPPhone(signup.CountryCode,signup.PhoneNumber,otp,subject)
	}
	if err != nil{
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse:=response.SuccessResponse{Message: constant.OTP_SENT}
	utils.SuccessMessageResponse(w, 200, successResponse)
}
