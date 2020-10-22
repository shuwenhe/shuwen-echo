package dao

import (
	"github.com/shuwenhe/shuwen-beego/models"
	"github.com/shuwenhe/shuwen-beego/utils"
)

func ArticlePage(pageNo int, pageSize int) ([]*models.Article, error) {
	articles := make([]*models.Article, 0, pageSize)
	sql := "select * from article limit ?,?"
	err := utils.Db.Select(&articles, sql, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func ArticleCount() (int, error) {
	var count int
	sql := "SELECT COUNT(id) as count FROM article"
	err := utils.Db.Get(&count, sql)
	if err != nil {
		utils.NewResult(utils.Fail, err.Error())
		return 0, err
	}
	return count, nil
}
