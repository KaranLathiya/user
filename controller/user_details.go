package controller

import (
	"net/http"
	error_handling "user/error"
	"user/middleware"
	"user/model/response"
	"user/utils"

	"github.com/go-chi/chi"
)

func (c *UserController) GetUsernameByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	id := chi.URLParam(r, "id")
	isBlocked, err := c.repo.IsBlocked(userID, id)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	if !isBlocked {
		username, err := c.repo.GetUsernameByID(id)
		if err != nil {
			error_handling.ErrorMessageResponse(w, err)
			return
		}
		usernameResponse := response.Username{Username: username}
		utils.SuccessMessageResponse(w, 200, usernameResponse)
		return
	}
	utils.SuccessMessageResponse(w, 200, response.Username{})
}

func (c *UserController) GetUserDetailsByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserCtxKey).(string)
	id := chi.URLParam(r, "id")
	isBlocked, err := c.repo.IsBlocked(userID, id)
	if err != nil {
		error_handling.ErrorMessageResponse(w, err)
		return
	}
	if !isBlocked {
		userDetails, err := c.repo.GetUserDetailsByID(id)
		if err != nil {
			error_handling.ErrorMessageResponse(w, err)
			return
		}
		utils.SuccessMessageResponse(w, 200, userDetails)
		return 
	}
	utils.SuccessMessageResponse(w, 200, response.UserDetails{})
}
