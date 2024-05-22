package main

import (
	"fmt"
	"net/http"
	"user/config"
	"user/controller"
	"user/db"
	"user/repository"
	"user/routes"

	_ "user/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			User-Service API
//	@version		1.0
//	@description	User service for registration/login of user. It allows to block/unblock other users and update their profiles.
//	@host		localhost:8000/
// @tag.name UserAuth
// @tag.description User signup, login, google login
// @tag.name UserDetails
// @tag.description User details update, fetch 
// @tag.name BlockActions
// @tag.description Block or unblock another user  
// @tag.name User
// @tag.description get all users 
// @tag.name PublicAPI
// @tag.description inter service apis
// @schemes http

// @securitydefinitions.apikey UserIDAuth
// @in header
// @name Auth-user

// @securitydefinitions.apikey jwtAuth
// @in header
// @name Authorization
func main() {
	err := config.LoadConfig("../config")
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %v", err))
	}
	db := db.Connect()
	defer db.Close()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	router := routes.InitializeRouter(controllers)
	router.Mount("/swagger/", httpSwagger.WrapHandler)
	fmt.Println("server started")
	http.ListenAndServe(":"+config.ConfigVal.Port, router)
}
