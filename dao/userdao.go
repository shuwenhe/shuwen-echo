package dao

import (
	"github.com/shuwenhe/shuwen-beego/models"
	"github.com/shuwenhe/shuwen-echo/db"
)

func Login(num string) (*models.User, error) {
	sql := "select * from users where num=?"
	rows, err := db.Db.Query(sql, num)
	user := &models.User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Num, &user.Password, &user.Phone)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
