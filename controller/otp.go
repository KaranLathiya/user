package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

func (c *UserController) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var verifyOTP request.VerifyOTP
	err := utils.BodyReadAndValidate(r.Body, &verifyOTP, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.VerifyOTP(verifyOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	go c.repo.DeleteOTPs(verifyOTP)
	var userID string
	if verifyOTP.EventType == constant.EVENT_TYPE_SIGNUP {
		userID, err = c.repo.CreateUser(verifyOTP)
	} else if verifyOTP.EventType == constant.EVENT_TYPE_LOGIN {
		userID, err = c.repo.GetUserID(verifyOTP.Email, verifyOTP.PhoneNumber, verifyOTP.CountryCode, verifyOTP.SignupMode)
	} else if verifyOTP.EventType == constant.EVENT_TYPE_GOOGLE_LOGIN {
		userID, err = c.repo.GetUserID(verifyOTP.Email, verifyOTP.PhoneNumber, verifyOTP.CountryCode, verifyOTP.SignupMode)
		if err != nil {
			userID, err = c.repo.CreateUser(verifyOTP)
		}
	}
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, response.UserID{UserID: userID})
}
