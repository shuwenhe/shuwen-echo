package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/dao"
	"github.com/shuwenhe/utils"
)

func Login(ctx echo.Context) error {
	num := ctx.FormValue("num")
	if num == "" {
		return ctx.JSON(utils.NewErrIpt("num is not null"))
	}
	user, err := dao.Login(num)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, user)
}
