package main

import (
	"fmt"
	"net/http"
	"os"
	"go-structure/config"
	"go-structure/controller"
	"go-structure/database"
	"go-structure/repository"
	"go-structure/routes"
)

func main() {
	_, err := config.LoadConfig("./config")
	if err != nil {
		panic(fmt.Sprintf("cannot load config: %v", err))
	}
	db := database.Connect()
	defer db.Close()
	repos := repository.InitRepositories(db)
	controllers := controller.InitControllers(repos)
	router := routes.IntializeRouter(controllers)
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
