package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/middleware"
	"user/model/request"
	"user/model/response"
	"user/utils"

	"github.com/go-chi/chi"
)

func (c *UserController) GetUserDetailsByUsername(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	username := chi.URLParam(r, "username")
	userDetails, err := c.repo.GetUserDetailsByUsername(username, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userDetails)
}

func (c *UserController) GetUserDetailsByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	id := chi.URLParam(r, "user-id")
	userDetails, err := c.repo.GetUserDetailsByID(id, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userDetails)
}

func (c *UserController) GetCurrentUserDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	userDetails, err := c.repo.GetCurrentUserDetailsByID(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userDetails)
}

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
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}

func (c *UserController) UpdateBasicDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	var updateUserNameDetails request.UpdateUserNameDetails
	err := utils.BodyReadAndValidate(r.Body, &updateUserNameDetails, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.UpdateBasicDetails(updateUserNameDetails, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.USER_DETAILS_UPDATED}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}
