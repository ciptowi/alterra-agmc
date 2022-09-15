package controllers

import (
	"dynamic/lib"
	"dynamic/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func SaveUser(c echo.Context) error {
	usr := new(models.User)
	res := new(lib.Response)
	c.Bind(usr)

	if models.CreateUser(usr) != nil {
		res.Message = "Create user Failed"
		return c.JSON(http.StatusBadRequest, "bad request")
	} else {
		res.Message = "Create user Successfully"
		res.Data = *usr
	}
	return c.JSON(http.StatusCreated, res)
}

func GetUsers(c echo.Context) error {
	res := new(lib.Response)

	user, err := models.GetAll()
	if err != nil {
		res.Message = "Failed"
		return c.JSON(http.StatusBadRequest, "bad request")
	} else {
		res.Message = "Success"
		res.Data = user
		return c.JSON(http.StatusOK, res)
	}
}

func GetUserById(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func UpdateUserById(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func DeleteUserById(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
