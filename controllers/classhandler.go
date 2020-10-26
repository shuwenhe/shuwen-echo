package controllers

import (
	"net/http"
	"strconv"

	"github.com/shuwenhe/shuwen-echo/models"
	"github.com/shuwenhe/utils"

	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/dao"
)

func Index(c echo.Context) error {
	classes, _ := dao.GetClasses()
	return c.JSON(http.StatusOK, classes)
}

func AddClass(ctx echo.Context) error {
	name := ctx.FormValue("name")
	des := ctx.FormValue("des")
	cls := &models.Class{
		Name: name,
		Des:  des,
	}
	if cls.Name == "" {
		return ctx.JSON(utils.NewFail("The name is not null"))
	}
	class, err := dao.GetClassByName(name)
	if err != nil {
		err2 := dao.AddClass(cls)
		if err2 != nil {
			return ctx.JSON(utils.NewFail("Add the class fail!", err.Error()))
		}
		return ctx.JSON(utils.NewSucc("Add the class success!", cls))
	} else {
		if class.Name == name {
			return ctx.JSON(utils.NewFail("The name is exist!"))
		}
	}
	return ctx.JSON(utils.NewSucc("Add class success!"))
}

func GetClassByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return ctx.JSON(utils.NewFail("Get id fail!", err.Error()))
	}
	class, err := dao.GetClassByID(id)
	if err != nil {
		return ctx.JSON(utils.NewFail("Get class fail!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("Get class success!", class))
}

func GetClasses(ctx echo.Context) error {
	classes, err := dao.GetClasses()
	if err != nil {
		return ctx.JSON(utils.NewFail("Get classes fail!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("Get the classes success!", classes))
}

func UpdateClass(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return ctx.JSON(utils.NewFail("Get the id is fail!", err.Error()))
	}
	name := ctx.FormValue("name")
	des := ctx.FormValue("des")
	class := &models.Class{
		ID:   id,
		Name: name,
		Des:  des,
	}
	err2 := dao.UpdateClass(class)
	if err2 != nil {
		return ctx.JSON(utils.NewFail("Get the class is fail!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("Update class success!", class))
}

func DeleteClass(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return ctx.JSON(utils.NewFail("Get the id is fail!"))
	}
	class, err := dao.GetClassByID(id)
	if err != nil {
		return ctx.JSON(utils.NewFail("The id is not exist,please try it again!"))
	}
	if class.ID > 0 {
		err2 := dao.DeleteClass(id)
		if err2 != nil {
			return ctx.JSON(utils.NewFail("Delete the id is fail!"))
		}
	}
	return ctx.JSON(utils.NewSucc("Delete class success!"))
}
