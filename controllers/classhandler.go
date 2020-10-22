package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/dao"
)

func Index(c echo.Context) error {
	classes, err := dao.GetClasses()
	fmt.Println(classes, err)
	fmt.Printf("%#v", classes)
	return c.JSON(http.StatusOK, classes)
}
