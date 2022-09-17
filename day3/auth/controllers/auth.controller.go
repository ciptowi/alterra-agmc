package controllers

import (
	"auth/helpers"
	"auth/lib"
	"auth/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type tokenString struct {
	AccsessToken string `json:"accsessToken"`
	Type         string `json:"type"`
}

var (
	JWT_SECRET         = []byte("secreet-keyy")
	JWT_EXP            = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

func LoginByEmailAndPassword(c echo.Context) error {
	usr := new(models.User)
	res := new(helpers.Response)

	my_data := echo.Map{}
	if err := c.Bind(&my_data); err != nil {
		return err
	}
	em := fmt.Sprintf("%v", my_data["email"])
	pss := fmt.Sprintf("%v", my_data["password"])

	u := lib.FindUserByEmail(em, usr)
	if u != nil {
		res.Message = "Email Not Found"
		return c.JSON(http.StatusFound, res)
	} else if helpers.CheckPasswordHash(pss, usr.Password) == false {
		res.Message = "Wrong Password"
		return c.JSON(http.StatusFound, res)
	}
	str := new(tokenString)
	token, _ := generateToken(usr)

	str.AccsessToken = token
	str.Type = "Bearer"
	res.Success = true
	res.Message = "Login Success"
	res.Data = str
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

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// type Claims struct {
// 	Id   string `json:"id"`
// 	Name string `json:"name"`
// 	jwt.StandardClaims
// }

// func getTokenString(authHeader string) (*string, error) {
// 	var token string
// 	if strings.Contains(authHeader, "Bearer") {
// 		token = strings.Replace(authHeader, "Bearer ", "", -1)
// 		return &token, nil
// 	}
// 	return nil, fmt.Errorf("authorization not found")
// }

// func CreateJWTClaims(user, *models.User) dto.JWTClaims {
// 	return dto.JWTClaims{
// 		UserID:     userID,
// 		Email:      email,
// 		RoleID:     roleID,
// 		DivisionID: divisionID,
// 		RegisteredClaims: jwt.st{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXP)),
// 		},
// 	}
// }

// func createJWTToken(claims Claims) (string, error) {
// 	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
// 	return token.SignedString([]byte(JWT_SECRET))
// }

// func parseJWTToken(authHeader string) (*Claims, error) {
// 	tokenString, err := getTokenString(authHeader)
// 	if err != nil {
// 		return nil, err
// 	}
// 	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != JWT_SIGNING_METHOD {
// 			return nil, fmt.Errorf("invalid signing method")
// 		}
// 		return JWT_SECRET, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		claimsStr, err := json.Marshal(claims)
// 		if err != nil {
// 			return nil, fmt.Errorf("error when marshalling token")
// 		}

// 		var customClaims Claims
// 		if err := json.Unmarshal(claimsStr, &customClaims); err != nil {
// 			return nil, fmt.Errorf("error when unmarshalling token")
// 		}

// 		return &customClaims, nil
// 	} else {
// 		return nil, fmt.Errorf("invalid token")
// 	}
// }
