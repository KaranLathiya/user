package controller

import (
	"net/http"
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
	if verifyOTP.EventType == "signup" {
		userID, err = c.repo.CreateUser(verifyOTP)
	} else if verifyOTP.EventType == "login" {
		userID, err = c.repo.GetUserID(verifyOTP)
	} else if verifyOTP.EventType == "google_login" {
		userID, err = c.repo.GetUserID(verifyOTP)
		if err != nil {
			userID, err = c.repo.CreateUser(verifyOTP)
		}
	}
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, response.UserID{UserID: userID})
}
