package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shuwenhe/shuwen-echo/controllers"
	"github.com/shuwenhe/shuwen-echo/controllers/test"
)

func Run() {
	e := echo.New()

	e.HideBanner = true

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

	api.POST("/class/addClass", controllers.AddClass)
	api.GET("/class/getClassByID", controllers.GetClassByID)
	api.GET("/class/getClasses", controllers.GetClasses)
	api.POST("/class/updateClass", controllers.UpdateClass)
	api.POST("/class/deleteClass", controllers.DeleteClass)

	api.GET("/article/getArticleByID", controllers.GetArticleByID)
	api.POST("/article/addArticle", controllers.AddArticle)

	e.Start(":8080")
}
