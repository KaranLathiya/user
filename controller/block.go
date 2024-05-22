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

// Block User example
//
// @tags BlockActions
// @Security UserIDAuth
//	@Summary		block another user
//	@Description	block another user to hide your details
//	@ID				user-block
//	@Accept			json
//	@Produce		json
// @Param request body request.BlockUser true "The input for user login"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/block/ [post]
func (c *UserController) BlockUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	var blockUser request.BlockUser
	err := utils.BodyReadAndValidate(r.Body, &blockUser, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.BlockUser(blockUser, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	successResponse := response.SuccessResponse{Message: constant.BLOCKED_USER}
	utils.SuccessMessageResponse(w, http.StatusOK, successResponse)
}

// Unblock User example
//
// @tags BlockActions
// @Security UserIDAuth
//	@Summary		unblock another user
//	@Description	unblock another user to show your details 
//	@ID				user-unblock
//	@Accept			json
//	@Produce		json
// @Param blocked path string true "blocked"
//	@Success		200		{object}	response.SuccessResponse "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/unblock/{blocked} [delete]
func (c *UserController) UnblockUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	blockedUser := request.BlockUser{
		BlockedUser: chi.URLParam(r, "blocked"),
	}
	err := utils.ValidateStruct(blockedUser, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	err = c.repo.UnblockUser(blockedUser, userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	// successResponse := response.SuccessResponse{Message: constant.UNBLOCKED_USER}
	// utils.SuccessMessageResponse(w, http.StatusOK, nil)
	w.WriteHeader(http.StatusNoContent)
}

// Block UserList example
//
// @tags BlockActions
// @Security UserIDAuth
//	@Summary		block userList
//	@Description	get list of all blocked users 
//	@ID				user-blockList
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]response.BlockUserDetails "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/block [get]
func (c *UserController) BlockedUserList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	blockedUserList, err := c.repo.BlockedUserList(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, blockedUserList)
}
