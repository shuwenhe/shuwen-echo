package controllers

import (
	"strconv"
	"time"

	"github.com/shuwenhe/shuwen-echo/models"

	"github.com/labstack/echo"
	"github.com/shuwenhe/shuwen-echo/dao"
	"github.com/shuwenhe/utils"
)

func GetArticleByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.FormValue("id"))
	if err != nil {
		return ctx.JSON(utils.NewFail("Get the id is fail!", err.Error()))
	}
	article, err2 := dao.GetArticleByID(id)
	if err2 != nil {
		return ctx.JSON(utils.NewFail("Get the article fail!", err.Error()))
	}
	article.Class, err = dao.GetClassByID(article.Cid)
	if err != nil {
		return ctx.JSON(utils.NewFail("Get the class fail!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("Get the article success!", article))
}

func AddArticle(ctx echo.Context) error {
	cid, err := strconv.Atoi(ctx.FormValue("cid"))
	if err != nil {
		return ctx.JSON(utils.NewFail("Get cid is fail!", err.Error()))
	}
	title := ctx.FormValue("title")
	author := ctx.FormValue("author")
	content := ctx.FormValue("content")
	hits, err2 := strconv.Atoi(ctx.FormValue("hits"))
	if err2 != nil {
		return ctx.JSON(utils.NewFail("Get cid is fail!", err.Error()))
	}
	a := &models.Article{
		Cid:     cid,
		Title:   title,
		Author:  author,
		Content: content,
		Hits:    hits,
		Ctime:   time.Now(),
	}
	err3 := dao.AddArticle(a)
	if err3 != nil {
		return ctx.JSON(utils.NewFail("Add article is fail!", err.Error()))
	}
	return ctx.JSON(utils.NewSucc("Add article is success!"))
}
