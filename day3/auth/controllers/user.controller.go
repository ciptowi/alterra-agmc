package controllers

import (
	"auth/lib"
	"auth/models"
	"auth/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	usr := new(models.User)
	res := new(lib.Response)

	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	}
	nm := fmt.Sprintf("%v", my_data["name"])
	em := fmt.Sprintf("%v", my_data["email"])
	pss := fmt.Sprintf("%v", my_data["password"])
	hash, err := utils.HashPassword(pss)
	if err != nil {
		return err
	}
	usr.Name = nm
	usr.Email = em
	usr.Password = hash
	lib.CreateUser(usr)
	res.Message = "Created"
	res.Data = usr
	return c.JSON(http.StatusCreated, res)
}

func GetUsers(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	users := new(models.Users)
	res := new(lib.Response)

	lib.FindUsers(users)
	res.Message = "Success"
	res.Data = users
	return c.JSON(http.StatusOK, res)
}

func GetUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	lib.FindUserById(id, usr)
	res.Message = "Success"
	res.Data = usr
	return c.JSON(http.StatusOK, res)
}

func UpdateUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	if err := c.Bind(&usr); err != nil {
		return err
	}
	lib.UpdateUserById(id, usr)
	res.Message = "Updated"
	return c.JSON(http.StatusOK, res)
}

func DeleteUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	lib.DeleteUserById(id, usr)
	res.Message = "Deleted"
	return c.JSON(http.StatusOK, res)
}
