package dao

import (
	"github.com/shuwenhe/shuwen-beego/models"
	"github.com/shuwenhe/shuwen-echo/db"
	"github.com/shuwenhe/shuwen-go-web/model"
)

func AddClass(cls *models.Class) error {
	sql := "insert into class (name,des) value(?,?)"
	_, err := db.Db.Exec(sql, cls.Name, cls.Des)
	if err != nil {
		return err
	}
	return nil
}

func UpdateClass(cls *models.Class) error {
	sql := "update class set name=?,des=? where id=?"
	_, err := db.Db.Exec(sql, cls.Name, cls.Des, cls.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetClasses() ([]*model.Class, error) {
	sql := "select * from class"
	classes := []*model.Class{}
	err := db.Db.Select(&classes, sql)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
