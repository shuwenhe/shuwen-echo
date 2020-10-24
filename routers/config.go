package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/models"
	"github.com/shuwenhe/utils"
)

func Filter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		if token == "" {
			return c.JSON(utils.NewFail("Please login first"))
		}
		dl := &models.Jwt{}
		j, err := jwt.ParseWithClaims(token, dl, func(token *jwt.Token) (interface{}, error) {
			return []byte("key"), nil
		})
		if err != nil {
			return c.JSON(utils.NewFail("Login information verification failed"))
		}
		if j.Valid {
			c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
			return next(c)
		} else {
			return c.JSON(utils.NewFail("Login information verification failed"))
		}
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
