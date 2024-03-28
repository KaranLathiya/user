package controller

import (
	"net/http"
	"go-structure/repository"
)

type UserController struct {
	repo repository.Repository
}

func InitControllers(repo repository.Repository) *UserController {
	return &UserController{repo: repo}
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	
	err := c.repo.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
}
