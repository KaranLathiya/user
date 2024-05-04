package controller

import (
	"encoding/json"
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

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

func (c *UserController) CreateOTPForDeleteOrganization(w http.ResponseWriter, r *http.Request) {
	err := utils.VerifyJWT(r.Header.Get("Authorization"), "Organization", "Organization")
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
	body, err := utils.CallAnotherService(jwtToken, constant.ORGANIZATION_SERVICE_BASE_URL+"internal/organization/"+verifyOTPForDeleteOrganization.OrganizationID, nil, http.MethodDelete)
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
