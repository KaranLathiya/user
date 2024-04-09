package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/middleware"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

func (c *UserController) UpdateUserPrivacy(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	var updateUserPrivacy request.UpdateUserPrivacy
	err := utils.BodyReadAndValidate(r.Body, &updateUserPrivacy, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.UpdateUserPrivacy(updateUserPrivacy, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.USER_DETAILS_UPDATED}
	utils.SuccessMessageResponse(w, 200, successResponse)
}

func (c *UserController) UpdateUserNameDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	var updateUserNameDetails request.UpdateUserNameDetails
	err := utils.BodyReadAndValidate(r.Body, &updateUserNameDetails, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.UpdateUserNameDetails(updateUserNameDetails, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.USER_DETAILS_UPDATED}
	utils.SuccessMessageResponse(w, 200, successResponse)
}
