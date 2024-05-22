package controller

import (
	"net/http"
	"strconv"
	error_handling "user/error"
	"user/middleware"
	"user/model/request"
	"user/utils"
)

// get userList example
//
// @tags User
// @Security UserIDAuth
//	@Summary		get userList
//	@Description	get userList using different sorting & sorting
//	@ID				user-list
//	@Accept			json
//	@Produce		json
// @Param     limit        query     int        false  "10"
// @Param     offset       query     int        false  "0"
// @Param     email        query     string     false  " "
// @Param     fullname     query     string     false  " "
// @Param     phonenumber  query     string     false  " "
// @Param     sorting      query     string     false  "pass the value asc or desc"
// @Param     filter       query     string     false  "pass the value fullname or date"
//	@Success		200		{object}	[]response.User "OK"
//	@Failure		400		{object}	error.CustomError	"Bad Request"
//	@Failure		401		{object}	error.CustomError	"Unauthorized"
//	@Failure		404		{object}	error.CustomError	"Not Found"
//	@Failure		405		{object}	error.CustomError	"Method Not Allowed"
//	@Failure		409		{object}	error.CustomError	"Conflict"
//	@Failure		500		{object}	error.CustomError	"Internal Server Error"
//	@Router			/users/ [get]
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
