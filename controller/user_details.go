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

// User details by username example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		user details
//	@Description	get another user details by username
//	@ID				user-details-by-username
//	@Accept			json
//	@Produce		json
// @Param  username path string true "username"
//	@Success		200		{object}	response.UserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/username/{username} [get]
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

// User details by userid example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		user details
//	@Description	get another user details by userID
//	@ID				user-details-by-userid
//	@Accept			json
//	@Produce		json
// @Param  user-id path string true "user-id"
//	@Success		200		{object}	response.UserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/id/{user-id} [get]
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

// Current user details example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		current user details
//	@Description	get current user details by userID
//	@ID				current-user-details
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.UserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/user/profile/ [get]
func (c *UserController) GetCurrentUserDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	userDetails, err := c.repo.GetCurrentUserDetailsByID(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userDetails)
}

// update user privacy example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		update user privacy
//	@Description	update user privacy
//	@ID				update-user-privacy
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.UserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/user/profile/privacy [put]
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

// update user basic details example
//
// @tags UserDetails
// @Security UserIDAuth
//	@Summary		update user basic details
//	@Description	update user basic details like firstname, lastname, username
//	@ID				update-user-basic
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	request.UpdateUserNameDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/user/profile/basic [put]
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
