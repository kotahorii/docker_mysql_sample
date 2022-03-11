package controller

import (
	"docker_sample/database"
	"docker_sample/schema"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []schema.User{}
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := schema.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := schema.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	user := schema.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&schema.User{}, id)
	return c.NoContent(http.StatusNoContent)
}
