package handler

import (
	"integration-test/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.SaveBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, bookJSON, rec.Body.String())
	}
}

func TestGetBooks(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.GetBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// func TestUpdateBookById(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPut, "/books/:id", strings.NewReader(updateBookJSON))
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("1")

// 	// Assertions
// 	if assert.NoError(t, controllers.UpdateBookById(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		// assert.Equal(t, updateBookJSON, rec.Body.String())
// 	}
// }

func TestDeleteBookById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.DeleteBookById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
