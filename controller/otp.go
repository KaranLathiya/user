package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

// VerifyOTP example
//
// @tags UserAuth
//	@Summary		verify otp
//	@Description	verify otp for login/signup 
//	@ID				verify-otp
//	@Accept			json
//	@Produce		json
// @Param request body request.VerifyOTP true "input for verify otp"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/otp/verify/ [post]
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
