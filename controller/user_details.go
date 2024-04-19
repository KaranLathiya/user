package controller

import (
	"fmt"
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
	id, err := c.repo.GetIDByUsername(username)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	userDetails, err := c.repo.GetUserDetailsByID(id, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, userDetails)
}

func (c *UserController) GetUserDetailsByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	id := chi.URLParam(r, "id")
	userDetails, err := c.repo.GetUserDetailsByID(id, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, userDetails)
}

func (c *UserController) GetCurrentUserDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	userDetails, err := c.repo.GetCurrentUserDetailsByID(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, userDetails)
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
	utils.SuccessMessageResponse(w, 200, successResponse)
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
	utils.SuccessMessageResponse(w, 200, successResponse)
}

func (c *UserController) GetUsersDetailsByIDs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("called")
	err := middleware.VerifyJWTToken(r.Header.Get("Authorization"), "User", "User details of organization")
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
	utils.SuccessMessageResponse(w, 200, usersDetails)
}
