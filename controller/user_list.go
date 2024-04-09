package controller

import (
	"net/http"
	"strconv"
	"strings"
	error_handling "user/error"
	"user/middleware"
	"user/model/request"
	"user/utils"
)

func (c *UserController) GetUserList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)

	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	orderBy := r.FormValue("orderby")
	order := r.FormValue("order")
	email := r.FormValue("email")
	fullname := r.FormValue("fullname")
	phone := r.FormValue("phone")

	//default limit
	if limit == 0 {
		limit = 10
	}

	userListParameter := request.UserListParameter{
		Limit:       limit,
		Offset:      offset,
		Order:       order,
		OrderBy:     orderBy,
		Email:       email,
		Fullname:    fullname,
		PhoneNumber: phone,
	}

	err := utils.ValidateStruct(userListParameter, nil)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}

	var where []string
	var filterArgsList []interface{}

	if fullname != "" {
		where = append(where, "fullname ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, r.FormValue("fullname"))
	}
	if email != "" {
		where = append(where, "email ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, r.FormValue("email"))
	}
	if phone != "" {
		where = append(where, "phone_number ILIKE '%' || ? || '%'")
		filterArgsList = append(filterArgsList, r.FormValue("phone"))
	}

	if orderBy == "" {
		orderBy = "fullname"
	} else if orderBy == "phone" {
		orderBy = "phone_number"
	} else if orderBy == "date" {
		orderBy = "created_at"
	}

	if order == "" {
		order = "asc"
	}

	blockedUserList, err := c.repo.BlockedUserList(userID)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	var blockedUserIDs []string
	for _, value := range blockedUserList {
		blockedUserIDs = append(blockedUserIDs, value.BlockedUser)
	}
	var blockedUserCondition string
	if len(blockedUserIDs) > 0 {
		blockedUserIDsString := "'" + strings.Join(blockedUserIDs, "' , '") + "'"
		blockedUserCondition = "id NOT IN (" + blockedUserIDsString + ")"
	}
	where = append(where, blockedUserCondition)

	userList, err := c.repo.GetUserList(userID, where, filterArgsList, orderBy, order, limit, offset, blockedUserIDs)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	utils.SuccessMessageResponse(w, 200, userList[0])
}
