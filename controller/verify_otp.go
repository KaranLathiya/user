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
	var addValidationRules map[string]string

	err := utils.BodyReadAndValidate(r.Body, &verifyOTP, addValidationRules)
	if err != nil {
		if verifyOTP.EventType == "signup" {
		addValidationRules = map[string]string{
			"firstname": "required|string|min_len:2|max_len:50",
			"lastname":  "required|string|min_len:2|max_len:50",
		}
	}
	err = utils.ValidateStruct(&verifyOTP, addValidationRules)
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	if verifyOTP.EventType == "signup" {
		addValidationRules = map[string]string{
			"Firstname": "required|string|min_len:2|max_len:50",
			"Lastname":  "required|string|min_len:2|max_len:50",
		}
	}
	err = utils.ValidateStruct(&verifyOTP, addValidationRules)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	// err = c.repo.VerifyOTP(verifyOTP)
	// if err != nil {
	// 	error_handling.ErrorMessageResponse(w, err)
	// 	return
	// }
	go c.repo.DeleteOTPs(verifyOTP)
	var userID string
	if verifyOTP.EventType == "signup" {
		userID, err = c.repo.UserCreate(verifyOTP)
		if err != nil {
			error_handling.ErrorMessageResponse(w, err)
			return
		}
	} else if verifyOTP.EventType == "login" {
		userID, err = c.repo.GetUserID(verifyOTP)
		if err != nil {
			error_handling.ErrorMessageResponse(w, err)
			return
		}
	} else if verifyOTP.EventType == "google_login" {
		userID, err = c.repo.GetUserID(verifyOTP)
		if err != nil {
			userID, err = c.repo.UserCreate(verifyOTP)
			if err != nil {
				error_handling.ErrorMessageResponse(w, err)
				return
			}
		}
	}
	userIDResponse := response.UserID{UserID: userID}
	utils.SuccessMessageResponse(w, 200, userIDResponse)
}
