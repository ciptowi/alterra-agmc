package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"static/lib"
	model "static/models"
	"strconv"
)

var (
	books = map[int]*model.Book{}
	seq   = 1
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetBooks(c echo.Context) error {
	res := new(lib.Response)

	res.Message = "Success"
	res.Data = books
	return c.JSON(http.StatusOK, res)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res := new(lib.Response)
	//for _, book := range books {
	//	if book.Id == id {
	//		res.Message = "Success"
	//		res.Data = book
	//	}
	//}

	if books[id] == nil {
		res.Message = "Not Found"
		return c.JSON(http.StatusNotFound, res)
	}
	res.Message = "Success"
	res.Data = books[id]
	return c.JSON(http.StatusOK, res)
}

func SaveBook(c echo.Context) error {
	//var newBook Book
	b := &model.Book{
		Id: seq,
	}
	res := new(lib.Response)

	err := c.Bind(b)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	//books = append(books, newBook)
	books[b.Id] = b
	seq++
	res.Data = b
	res.Message = "Create Success"
	return c.JSON(http.StatusCreated, res)
}

func UpdateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := new(model.Book)
	res := new(lib.Response)

	err := c.Bind(&b)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	//for _, book := range books {
	//	if book.Id == id {
	//		book.Id = b.Id
	//		book.Title = b.Title
	//		book.Writer = b.Writer
	//		book.Isbn = b.Isbn
	//	}
	//}
	books[id].Id = b.Id
	books[id].Title = b.Title
	books[id].Writer = b.Writer
	books[id].Isbn = b.Isbn
	res.Data = b
	res.Message = "Update Success"
	return c.JSON(http.StatusCreated, res)
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	res := new(lib.Response)

	delete(books, id)
	//return c.NoContent(http.StatusNoContent)
	res.Message = "Delete Success"
	return c.JSON(http.StatusNoContent, res)
}
