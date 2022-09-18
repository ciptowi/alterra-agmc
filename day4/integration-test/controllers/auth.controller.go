package controllers

import (
	"fmt"
	"integration-test/helpers"
	"integration-test/lib"
	"integration-test/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	JWT_EXP            = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

func LoginByEmailAndPassword(c echo.Context) error {
	usr := new(models.User)
	res := new(helpers.ResponseLogin)

	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	}
	em := fmt.Sprintf("%v", my_data["email"])
	pss := fmt.Sprintf("%v", my_data["password"])

	u := lib.FindUserByEmail(em, usr)
	if u != nil {
		res.Message = "Email Not Found"
		return c.JSON(http.StatusForbidden, res)
	} else if helpers.CheckPasswordHash(pss, usr.Password) == false {
		res.Message = "Wrong Password"
		return c.JSON(http.StatusForbidden, res)
	}
	token, _ := generateToken(usr)

	res.Success = true
	res.AccessToken = token
	res.TokenType = "Bearer"
	res.Message = "Login Success"
	res.Data = map[string]interface{}{
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
	}
	return c.JSON(http.StatusOK, res)
}

func generateToken(user *models.User) (string, error) {
	// sid := strconv.FormatUint(uint64(user.ID), 10)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(helpers.JWT_SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
