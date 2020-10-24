package routers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/models"
	"github.com/shuwenhe/shuwen-echo/util"
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
			return c.JSON(utils.NewFail("Login information verification failed!"))
		}
		if j.Valid {
			oldToken, ok := util.Get(dl.ID)
			if !ok {
				util.Set(dl.ID, token)
				c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
				return next(c)
			}
			if token != oldToken {
				return c.JSON(utils.NewFail("Your account logged in elsewhere!"))
			}
			c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
			return next(c)
		} else {
			return c.JSON(utils.NewFail("Login information verification failed"))
		}
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}
