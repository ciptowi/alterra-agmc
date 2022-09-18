package controllers

import (
	"integration-test/helpers"
	"integration-test/lib"
	"integration-test/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveBook(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	b := new(models.Book)
	res := new(helpers.Response)
	if err := c.Bind(&b); err != nil {
		return err
	}

	lib.CreateBook(b)
	res.Success = true
	res.Message = "Created"
	res.Data = b
	return c.JSON(http.StatusCreated, res)
}

func GetBooks(c echo.Context) error {
	books := new(models.Books)
	res := new(helpers.Response)

	lib.FindBooks(books)
	res.Success = true
	res.Message = "Success"
	res.Data = books
	return c.JSON(http.StatusOK, res)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(helpers.Response)

	lib.FindBookById(id, b)
	res.Success = true
	res.Message = "Success"
	res.Data = b
	return c.JSON(http.StatusOK, res)
}

func UpdateBookById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(helpers.Response)

	if err := c.Bind(&b); err != nil {
		return err
	}
	lib.UpdateBookById(id, b)
	res.Success = true
	res.Message = "Updated"
	return c.JSON(http.StatusOK, res)
}

func DeleteBookById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(helpers.Response)

	lib.DeleteBookById(id, b)
	res.Success = true
	res.Message = "Deleted"
	return c.JSON(http.StatusOK, res)
}
