package dao

import (
	db "github.com/shuwenhe/shuwen-echo/database"
	"github.com/shuwenhe/shuwen-echo/models"
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

func GetArticleByID(id int) (*models.Article, error) {
	sql := "select * from article where id = ?"
	article := &models.Article{}
	row := db.Db.QueryRow(sql, id)
	err := row.Scan(&article.ID, &article.Cid, &article.Title, &article.Author, &article.Content, &article.Hits, &article.Ctime)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func AddArticle(a *models.Article) error {
	sql := "insert into article (id,cid,title,author,content,hits,ctime) values(?,?,?,?,?,?,?)"
	_, err := db.Db.Exec(sql, a.ID, a.Cid, a.Title, a.Author, a.Content, a.Hits, a.Ctime)
	if err != nil {
		return err
	}
	return nil
}
