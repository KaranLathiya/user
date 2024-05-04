package controller

import (
	"net/http"
	"strconv"
	error_handling "user/error"
	"user/middleware"
	"user/model/request"
	"user/utils"
)

func (c *UserController) GetUserList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)

	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	filter := r.FormValue("filter")
	sorting := r.FormValue("sorting")
	email := r.FormValue("email")
	fullname := r.FormValue("fullname")
	phoneNumber := r.FormValue("phoneNumber")

	//default limit
	if limit == 0 {
		limit = 10
	}

	userListParameter := request.UserListParameter{
		Limit:       limit,
		Offset:      offset,
		Filter:      filter,
		Sorting:     sorting,
		Email:       &email,
		Fullname:    &fullname,
		PhoneNumber: &phoneNumber,
	}

	err := utils.ValidateStruct(userListParameter, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}

	userList, err := c.repo.GetUserList(userID, userListParameter)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, http.StatusOK, userList)
}
