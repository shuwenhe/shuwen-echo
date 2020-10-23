package dao

import (
	"github.com/shuwenhe/shuwen-beego/models"
	db "github.com/shuwenhe/shuwen-echo/database"
)

func ArticlePage(pageNo int, pageSize int) ([]*models.Article, error) {
	articles := make([]*models.Article, 0, pageSize)
	sql := "select * from article limit ?,?"
	err := db.Db.Select(&articles, sql, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func ArticleCount() (int, error) {
	var count int
	sql := "SELECT COUNT(id) as count FROM article"
	err := db.Db.Get(&count, sql)
	if err != nil {
		return 0, err
	}
	return count, nil
}
