package dao

import (
	db "github.com/shuwenhe/shuwen-echo/database"
	"github.com/shuwenhe/shuwen-echo/models"
)

func AddClass(class *models.Class) error {
	sql := "insert into class (id,name,des) value(?,?,?)"
	_, err := db.Db.Exec(sql, &class.ID, &class.Name, &class.Des)
	if err != nil {
		return err
	}
	return nil
}

func GetClasses() ([]*models.Class, error) {
	sql := "select * from class"
	classes := []*models.Class{}
	err := db.Db.Select(&classes, sql)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func GetClassByName(name string) (*models.Class, error) {
	sql := "select * from class where name=?"
	row := db.Db.QueryRow(sql, name)
	class := &models.Class{}
	err := row.Scan(&class.ID, &class.Name, &class.Des)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func GetClassByID(id int) (*models.Class, error) {
	sql := "select * from class where id=?"
	row := db.Db.QueryRow(sql, id)
	class := &models.Class{}
	err := row.Scan(&class.ID, &class.Name, &class.Des)
	if err != nil {
		return nil, err
	}
	return class, nil
}

func UpdateClass(cls *models.Class) error {
	sql := "update class set name=?,des=? where id=?"
	_, err := db.Db.Exec(sql, cls.Name, cls.Des, cls.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteClass(id int) error {
	sql := "delete from class where id=?"
	_, err := db.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
