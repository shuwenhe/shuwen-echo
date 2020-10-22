package router

import (
	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/controllers"
	"github.com/shuwenhe/shuwen-echo/controllers/test"
)

func Run() {
	app := echo.New()
	app.GET("/api/class/all", controllers.Index)
	app.GET("/api/test/index2", test.Index2)
	app.Start(":8080")
}
