package routers

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

	api := e.Group("/api", Filter)
	e.POST("/login", controllers.Login)
	e.POST("/addUser", controllers.AddUser)
	api.GET("/user/getUsers", controllers.GetUsers)
	api.POST("/user/updateUser", controllers.UpdateUser)
	api.GET("/user/del:id", controllers.DeleteUserByID)

	e.Start(":8080")
}
