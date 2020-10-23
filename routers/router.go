package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shuwenhe/shuwen-echo/controllers"
	"github.com/shuwenhe/shuwen-echo/controllers/test"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/class/all", controllers.Index)
	e.GET("/api/test/index2", test.Index2)
	e.GET("/user/login", controllers.Login)

	e.Start(":8080")
}
