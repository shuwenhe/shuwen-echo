package main

import (
	"github.com/shuwenhe/shuwen-echo/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/test", controller.Index)
	e.Start(":8080")
}
