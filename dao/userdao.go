package dao

import (
	"fmt"

	db "github.com/shuwenhe/shuwen-echo/database"
	"github.com/shuwenhe/shuwen-echo/models"
)

func Login(name string) (*models.User, error) {
	sql := "select * from users where name=?"
	rows, err := db.Db.Query(sql, name)
	user := &models.User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Password, &user.Phone)
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(user *models.User) error {
	sql := "insert into users(id,name,password,phone) values(?,?,?,?)"
	_, err := db.Db.Exec(sql, &user.ID, &user.Name, &user.Password, &user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int) error {
	sql := "delete from users where id=?"
	_, err := db.Db.Exec(sql, id)
	fmt.Println("***")
	if err != nil {
		fmt.Println("***")
		return err
	}
	return nil
}

func GetUserByID(id int64) (*models.User, error) {
	sql := "select * from users where id = ?"
	row := db.Db.QueryRow(sql, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Phone)
	if err != nil {
		return nil, err
	}
	fmt.Println("user = ", user)
	return user, nil
}
