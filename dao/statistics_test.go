package dao

import (
	"fmt"
	"testing"
)

func testStatisticsAll(t *testing.T) {
	articles, _ := StatisticsAll()
	for _, mod := range articles {
		fmt.Println("mod = ", mod)
	}
}
