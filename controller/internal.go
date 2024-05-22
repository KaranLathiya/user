package controller

import (
	"encoding/json"
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"

	"github.com/go-chi/chi"
)


// Users details by IDs example
//
// @tags PublicAPI
// @Security jwtAuth
//	@Summary		get users details
//	@Description	get users details by userIDs
//	@ID				users-details-by-userIDs
//	@Accept			json
//	@Produce		json
// @Param request body request.UserIDs true "input for get users details"
//	@Success		200		{object}	map[string]response.UserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/internal/users/details/ [post]
func (c *UserController) GetUsersDetailsByIDs(w http.ResponseWriter, r *http.Request) {
	err := utils.VerifyJWT(r.Header.Get("Authorization"), "User", "User")
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	var userIDs request.UserIDs
	err = utils.BodyReadAndValidate(r.Body, &userIDs, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	usersDetails, err := c.repo.GetUsersDetailsByIDs(userIDs.UserIDs)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, usersDetails)
}

// Create otp for deleting organization example
//
// @tags PublicAPI
// @Security jwtAuth
//	@Summary		get otp for deleting organization
//	@Description	get otp for deleting organization in owner's registered email/phonenumber 
//	@ID				otp-for-delete-organization
//	@Accept			json
//	@Produce		json
// @Param request body request.CreateOTPForDeleteOrganization true "input for get otp for deleting organization"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/internal/user/organization/delete/otp/ [post]
func (c *UserController) CreateOTPForDeleteOrganization(w http.ResponseWriter, r *http.Request) {
	err := utils.VerifyJWT(r.Header.Get("Authorization"), "User", "User")
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	var createOTPForDeleteOrganization request.CreateOTPForDeleteOrganization
	err = utils.BodyReadAndValidate(r.Body, &createOTPForDeleteOrganization, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	userDetails, err := c.repo.GetCurrentUserDetailsByID(createOTPForDeleteOrganization.OwnerID)
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
	loginType := constant.LOGIN_TYPE_EMAIL
	if userDetails.Email == nil {
		loginType = constant.LOGIN_TYPE_PHONE_NUMBER
	}
	storeOTP := request.StoreOTP{
		Email:          userDetails.Email,
		PhoneNumber:    userDetails.PhoneNumber,
		CountryCode:    userDetails.CountryCode,
		EventType:      constant.EVENT_TYPE_ORGANIZATION_DELETE,
		LoginType:      loginType,
		HashedOTP:      hashedOTP,
		OrganizationID: &createOTPForDeleteOrganization.OrganizationID,
	}
	err = c.repo.StoreOTP(storeOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	subject := "OTP is for " + createOTPForDeleteOrganization.Name + " organization delete"
	if loginType == constant.LOGIN_TYPE_EMAIL {
		go utils.SendOTPInEmail(*userDetails.Email, otp, subject)
	} else {
		go utils.SendOTPInPhoneNumber(*userDetails.CountryCode, *userDetails.PhoneNumber, otp, subject)
	}
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.OTP_SENT}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}

// verify otp for deleting organization example
//
// @tags PublicAPI
//	@Summary		verify otp for deleting organization
//	@Description	verify otp for deleting organization
//	@ID				verify-otp-for-delete-organization
//	@Accept			json
//	@Produce		json
// @Param request body request.VerifyOTPForDeleteOrganization true "input for verify otp for deleting organization"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/internal/user/organization/delete/otp/verify [post]
func (c *UserController) VerifyOTPForDeleteOrganization(w http.ResponseWriter, r *http.Request) {
	var verifyOTPForDeleteOrganization request.VerifyOTPForDeleteOrganization
	err := utils.BodyReadAndValidate(r.Body, &verifyOTPForDeleteOrganization, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	verifyOTP := request.VerifyOTP{
		EventType:      verifyOTPForDeleteOrganization.EventType,
		OTP:            verifyOTPForDeleteOrganization.OTP,
		OrganizationID: &verifyOTPForDeleteOrganization.OrganizationID,
	}
	err = c.repo.VerifyOTP(verifyOTP)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	go c.repo.DeleteOTPs(verifyOTP)
	jwtToken, err := utils.CreateJWT("Organization", "Organization")
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	// Organization service call for delete organization
	body, err := utils.CallHttpService(jwtToken, constant.ORGANIZATION_SERVICE_BASE_URL+"internal/organization/"+verifyOTPForDeleteOrganization.OrganizationID, nil, http.MethodDelete)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	var successResponse response.SuccessResponse
	err = json.Unmarshal(body, &successResponse)
	if err != nil {
		error_handling.ErrorMessageResponse(w, error_handling.UnmarshalError)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}

// get organization creator details example
//
// @tags PublicAPI
// @Security jwtAuth
//	@Summary		organization creator details
//	@Description	organization creator details by ID
//	@ID				organization-creator-details
//	@Accept			json
//	@Produce		json
// @Param  user-id path string true "user-id"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/internal/users/details/{user-id} [get]
func (c *UserController) GetOrganizationCreatorDetailsByID(w http.ResponseWriter, r *http.Request) {
	err := utils.VerifyJWT(r.Header.Get("Authorization"), "User", "User")
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	userID := request.UserID{
		UserID: chi.URLParam(r, "user-id"),
	}
	userDetails, err := c.repo.GetCurrentUserDetailsByID(userID.UserID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userDetails)
}
