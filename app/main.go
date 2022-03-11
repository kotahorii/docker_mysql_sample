package main

import (
	"docker_sample/controller"
	"docker_sample/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}
