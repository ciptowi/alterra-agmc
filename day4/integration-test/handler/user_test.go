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

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.SaveUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// func TestUpdateUserById(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(updateUserJSON))
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetParamNames("id")
// 	c.SetParamValues("1")

// 	// Assertions
// 	if assert.NoError(t, controllers.UpdateUserById(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		// assert.Equal(t, updateuserJSON, rec.Body.String())
// 	}
// }

func TestDeleteUserById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, controllers.DeleteUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
