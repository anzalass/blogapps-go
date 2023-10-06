package main

import (
	// "echodatabase/controller"
	"echodatabase/model"
	"echodatabase/routes"
	// "github.com/labstack/echo/v4"
)

func main() {

	model.InitModel()
	model.Migrate()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000").Error())

}
