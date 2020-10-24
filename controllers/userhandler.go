package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/dao"
	"github.com/shuwenhe/shuwen-echo/models"
	"github.com/shuwenhe/shuwen-echo/util"
	"github.com/shuwenhe/utils"
)

func Login(ctx echo.Context) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.JSON(utils.NewErrIpt("Name is not null"))
	}
	user, err := dao.Login(name)
	if err != nil {
		return ctx.JSON(utils.NewFail("Please check if the username is correct!"))
	}
	password := ctx.FormValue("password")
	if password != user.Password {
		return ctx.JSON(utils.NewFail("Password error!"))
	}
	data := models.Jwt{
		ID:   user.ID,
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := jwts.SignedString([]byte(`key`))
	util.Set(user.ID, token)
	if err != nil {
		return ctx.JSON(utils.NewFail("System error, please login again!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("token is success!", token))
}

func AddUser(ctx echo.Context) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.JSON(utils.NewErrIpt("name is not null"))
	}
	user, err := dao.Login(name)
	if err != nil {
		return ctx.JSON(utils.NewFail("Please check if the username is correct!"))
	}
	if user.ID > 0 {
		return ctx.JSON(utils.NewFail("User is exist,please try it again!"))
	}
	password := ctx.FormValue("password")
	phone := ctx.FormValue("phone")
	user2 := &models.User{
		Name:     name,
		Password: password,
		Phone:    phone,
	}
	dao.AddUser(user2)
	return ctx.JSON(utils.NewSucc("token is success!", user2))
}

func DeleteUserByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println("id=***=", id)
	if err != nil {
		return ctx.JSON(utils.NewFail("Input data is error", err.Error()))
	}
	err2 := dao.DeleteUserByID(id)
	if err2 != nil {
		return ctx.JSON(utils.NewFail("Delete user fail"))
	}
	return ctx.JSON(utils.NewSucc("Delete user success"))
}

func GetUsers(ctx echo.Context) error {
	users, err := dao.GetUsers()
	if err != nil {
		return ctx.JSON(utils.NewFail("Get the users fail!"))
	}
	return ctx.JSON(utils.NewSucc("Get the users success!", users))
}
