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
	utils.SuccessMessageResponse(w, 200, successResponse)
}

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
	successResponse := response.SuccessResponse{Message: constant.UNBLOCKED_USER}
	utils.SuccessMessageResponse(w, 200, successResponse)
}

func (c *UserController) BlockUserList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	blockedUserList, err := c.repo.BlockedUserList(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, blockedUserList)
}
