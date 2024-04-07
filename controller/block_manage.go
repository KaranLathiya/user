package controller

import (
	"net/http"
	"user/constant"
	error_handling "user/error"
	"user/model/request"
	"user/model/response"
	"user/utils"
)

func (c *UserController) BlockUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Authorization")
	if userID == "" {
		error_handling.ErrorMessageResponse(w, error_handling.HeaderdataMisssing)
		return
	}
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
	userID := r.Header.Get("Authorization")
	if userID == "" {
		error_handling.ErrorMessageResponse(w, error_handling.HeaderdataMisssing)
		return
	}
	var blockedUser request.BlockUser
	err := utils.BodyReadAndValidate(r.Body, &blockedUser, nil)
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
	userID := r.Header.Get("Authorization")
	if userID == "" {
		error_handling.ErrorMessageResponse(w, error_handling.HeaderdataMisssing)
		return
	}
	blockedUserList,err := c.repo.BlockedUserList(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, blockedUserList)
}