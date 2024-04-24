package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var login request.Login
	err := utils.BodyReadAndValidate(r.Body, &login, nil)
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
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.BcryptError)
		return
	}
	storeOTP := request.StoreOTP{
		Email:       login.Email,
		PhoneNumber: login.PhoneNumber,
		CountryCode: login.CountryCode,
		EventType:   constant.EVENT_TYPE_LOGIN,
		LoginType:   login.LoginType,
		HashedOTP:   hashedOTP,
	}
	err = c.repo.StoreOTP(storeOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP for login "
	if login.LoginType == constant.LOGIN_TYPE_EMAIL {
		go utils.SendOTPInEmail(*login.Email, otp, subject)
	} else {
		go utils.SendOTPInPhoneNumber(*login.CountryCode, *login.PhoneNumber, otp, subject)
	}
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.OTP_SENT}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}
