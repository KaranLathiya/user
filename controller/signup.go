package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

// User Signup example
//
// @tags UserAuth
//	@Summary		user signup
//	@Description	new user signup using email/phoneNumber 
//	@ID				user-signup
//	@Accept			json
//	@Produce		json
// @Param request body request.Signup true "input for user signup"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/auth/signup/ [post]
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
		EventType:   constant.EVENT_TYPE_SIGNUP,
		LoginType:   signup.LoginType,
		HashedOTP:   hashedOTP,
	}
	err = c.repo.StoreOTP(storeOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP for signup: "
	if signup.LoginType == constant.LOGIN_TYPE_EMAIL {
		go utils.SendOTPInEmail(*signup.Email, otp, subject)
	} else {
		go utils.SendOTPInPhoneNumber(*signup.CountryCode, *signup.PhoneNumber, otp, subject)
	}
	successResponse := response.SuccessResponse{Message: constant.OTP_SENT}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}
