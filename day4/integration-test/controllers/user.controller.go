package controllers

import (
	"fmt"
	"integration-test/helpers"
	"integration-test/lib"
	"integration-test/models"
	"integration-test/restrict"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	usr := new(models.User)
	res := new(helpers.Response)

	body := echo.Map{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	nm := fmt.Sprintf("%v", body["name"])
	eml := fmt.Sprintf("%v", body["email"])
	pss := fmt.Sprintf("%v", body["password"])
	hash, err := helpers.HashPassword(pss)
	if err != nil {
		return err
	}
	if restrict.ChecklUserExisting(eml, usr) {
		res.Message = "Email already exists!"
		return c.JSON(http.StatusCreated, res)
	}
	usr.Name = nm
	usr.Email = eml
	usr.Password = hash
	lib.CreateUser(usr)
	res.Success = true
	res.Message = "Created"
	res.Data = usr
	return c.JSON(http.StatusCreated, res)
}

func GetUsers(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	users := new(models.Users)
	res := new(helpers.Response)

	lib.FindUsers(users)
	res.Success = true
	res.Message = "Success"
	res.Data = users
	return c.JSON(http.StatusOK, res)
}

func GetUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(helpers.Response)

	lib.FindUserById(id, usr)
	res.Success = true
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"id":        usr.ID,
		"name":      usr.Name,
		"email":     usr.Email,
		"createdAt": usr.CreatedAt,
		"updatedAt": usr.UpdatedAt,
		"deletedAt": usr.DeletedAt,
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(helpers.Response)

	if err := c.Bind(&usr); err != nil {
		return err
	}
	lib.UpdateUserById(id, usr)
	res.Success = true
	res.Message = "Updated"
	return c.JSON(http.StatusOK, res)
}

func DeleteUserById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(helpers.Response)

	lib.DeleteUserById(id, usr)
	res.Success = true
	res.Message = "Deleted"
	return c.JSON(http.StatusOK, res)
}
