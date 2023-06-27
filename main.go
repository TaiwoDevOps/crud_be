package main

import (
	"net/http"

	"taiwodevops.com/crud_be/config"
	"taiwodevops.com/crud_be/controllers"
	"taiwodevops.com/crud_be/helper"
	"taiwodevops.com/crud_be/repository"
	"taiwodevops.com/crud_be/router"
	"taiwodevops.com/crud_be/services"
)

func main() {

	// init DB
	db := config.DatabaseConnection()

	// use db instance to init repository
	repository := repository.NewBookRepositoryImpl(db)

	// use repository instance to init services
	services := services.NewBookServiceImpl(repository)

	// use services instance to init controllers

	controllers := controllers.NewBookController(services)

	// use controllers instance to init routes
	routes := router.NewRouter(controllers)

	// start server once all instances are init
	server := http.ListenAndServe("localhost:8888", routes)

	helper.PanicIfError(server)

}
