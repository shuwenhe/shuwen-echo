package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/shuwenhe/shuwen-echo/models"
)

func testArticlePage(t *testing.T) {
	articles, _ := ArticlePage(1, 1)
	for _, article := range articles {
		fmt.Println("article=", article)
	}
}

func testArticleCount(t *testing.T) {
	count, _ := ArticleCount()
	fmt.Println("count = ", count)
}

func testGetArticleByID(t *testing.T) {
	id := 1
	article, _ := GetArticleByID(id)
	fmt.Println("article = ", article)
}

func TestAddArticle(t *testing.T) {
	cid := 5
	title := "shuwen-sql"
	author := "ShuwenHe"
	content := "sql"
	hits := 20
	ctime := time.Now()
	a := &models.Article{
		Cid:     cid,
		Title:   title,
		Author:  author,
		Content: content,
		Hits:    hits,
		Ctime:   ctime,
	}
	AddArticle(a)
}
