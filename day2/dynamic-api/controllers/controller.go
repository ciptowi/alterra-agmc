package controllers

import (
	"dynamic/lib"
	"dynamic/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func SaveUser(c echo.Context) error {
	usr := new(models.User)
	res := new(lib.Response)
	if e := c.Bind(usr); e != nil {
		return e
	}

	if models.CreateUser(&usr) != nil {
		res.Message = "Create user Failed"
		return c.JSON(http.StatusBadRequest, "bad request")
	} else {
		res.Message = "Create user Successfully"
		res.Data = *usr
	}
	return c.JSON(http.StatusCreated, res)
}

func GetUsers(c echo.Context) error {
	usr := new(models.User)
	res := new(lib.Response)

	models.GetAll()
	//	res.Message = "Failed"
	//	return c.JSON(http.StatusBadRequest, "bad request")
	//} else {
	res.Message = "Success"
	res.Data = usr
	return c.JSON(http.StatusOK, res)
	//}
}

func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	models.GetOneById(id)
	//	res.Message = "Failed"
	//	return c.JSON(http.StatusBadRequest, "bad request")
	//} else {
	res.Message = "Success"
	res.Data = usr
	return c.JSON(http.StatusOK, res)
	//}
}

func UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	if models.UpdateUserById(id) != nil {
		res.Message = "Failed"
		return c.JSON(http.StatusBadRequest, "bad request")
	} else {
		res.Message = "Success"
		res.Data = usr
		return c.JSON(http.StatusOK, res)
	}
}

func DeleteUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	usr := new(models.User)
	res := new(lib.Response)

	if models.DeleteUserById(id) != nil {
		res.Message = "Failed"
		return c.JSON(http.StatusBadRequest, "bad request")
	} else {
		res.Message = "Success"
		res.Data = usr
		return c.JSON(http.StatusOK, res)
	}
}
