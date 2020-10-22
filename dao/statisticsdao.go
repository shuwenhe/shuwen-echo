package dao

import (
	"fmt"

	"github.com/shuwenhe/shuwen-echo/models"
	"github.com/shuwenhe/shuwen-echo/utils"
)

func StatisticsAll() ([]*models.Statistics, error) {
	Articles := []*models.Statistics{}
	sql := "select count(id) as count,DATE_FORMAT(ctime,'%Y-%m-%d') ctime from article group by date_format(ctime,'%Y-%m-%d')"
	err := utils.Db.Select(&Articles, sql)
	if err != nil {
		fmt.Println("err = ", err)
		return nil, err
	}
	return Articles, nil
}
