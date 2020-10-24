package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-echo/models"
)

func testLogin(t *testing.T) {
	num := "1"
	user, _ := Login(num)
	fmt.Println("user=", user)
}

func testAddUser(t *testing.T) {
	user := &models.User{
		Name:     "Test",
		Password: "123456",
		Phone:    "15010729356",
	}
	AddUser(user)
}

func testDeleteUser(t *testing.T) {
	DeleteUserByID(1)
}

func testGetUserByID(t *testing.T) {
	user, _ := GetUserByID(1)
	fmt.Println("user = ", user)
}
