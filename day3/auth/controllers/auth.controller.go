package controllers

import (
	"auth/lib"
	"auth/middleware"
	"auth/models"
	"auth/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type tokenString struct {
	accsessToken string
	test         string
}

func LoginByEmailAndPassword(c echo.Context) error {
	usr := new(models.User)
	res := new(lib.Response)

	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	}
	em := fmt.Sprintf("%v", my_data["email"])
	pss := fmt.Sprintf("%v", my_data["password"])

	u := lib.FindUserByEmail(em, usr)
	if u != nil {
		res.Success = false
		res.Message = "Email Not Found"
		return c.JSON(http.StatusFound, res)
	} else if utils.CheckPasswordHash(pss, usr.Password) == false {
		res.Success = false
		res.Message = "Wrong Password"
		return c.JSON(http.StatusFound, res)
	}
	str := new(tokenString)
	token, _ := middleware.GenerateToken(usr)

	str.accsessToken = token
	str.test = "test"
	res.Success = true
	res.Message = "Login Success"
	res.Data = str
	return c.JSON(http.StatusOK, res)
}
