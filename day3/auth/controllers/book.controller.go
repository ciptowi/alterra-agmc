package controllers

import (
	"auth/lib"
	"auth/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveBook(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	b := new(models.Book)
	res := new(lib.Response)
	if err := c.Bind(&b); err != nil {
		return err
	}

	lib.CreateBook(b)
	res.Message = "Created"
	res.Data = b
	return c.JSON(http.StatusCreated, res)
}

func GetBooks(c echo.Context) error {
	books := new(models.Books)
	res := new(lib.Response)

	lib.FindBooks(books)
	res.Message = "Success"
	res.Data = books
	return c.JSON(http.StatusOK, res)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(lib.Response)

	lib.FindBookById(id, b)
	res.Message = "Success"
	res.Data = b
	return c.JSON(http.StatusOK, res)
}

func UpdateBookById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(lib.Response)

	if err := c.Bind(&b); err != nil {
		return err
	}
	lib.UpdateBookById(id, b)
	res.Message = "Updated"
	return c.JSON(http.StatusOK, res)
}

func DeleteBookById(c echo.Context) error {
	// authHeader := c.Request().Header.Get("Authorization")
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(models.Book)
	res := new(lib.Response)

	lib.DeleteBookById(id, b)
	res.Message = "Deleted"
	return c.JSON(http.StatusOK, res)
}
