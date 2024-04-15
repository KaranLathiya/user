package main

import (
	"fmt"
	"net/http"
	"user/config"
	"user/controller"
	"user/database"
	"user/repository"
	"user/routes"
)

func main() {
	err := config.LoadConfig("../config")
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %v", err))
	}
	db := database.Connect()
	defer db.Close()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	router := routes.InitializeRouter(controllers)
	http.Handle("/", router)
	fmt.Println("server started")
	http.ListenAndServe(":"+config.ConfigVal.Port, nil)
}
