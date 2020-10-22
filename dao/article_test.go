package dao

import (
	"fmt"
	"testing"
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
