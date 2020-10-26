package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-echo/models"
)

func testAddClass(t *testing.T) {
	class := &models.Class{
		Name: "shuwen3",
		Des:  "mathematics",
	}
	AddClass(class)
}

func testGetClasses(t *testing.T) {
	classes, _ := GetClasses()
	for _, class := range classes {
		fmt.Println("class = ", class)
	}
}

func testGetClassByName(t *testing.T) {
	name := "shuwen"
	class, _ := GetClassByName(name)
	fmt.Println("class = ", class)
}

func testGetClassByID(t *testing.T) {
	id := 66
	class, _ := GetClassByID(id)
	fmt.Println("class = ", class)
}

func testUpdateClass(t *testing.T) {
	cls := &models.Class{
		ID:   67,
		Name: "shuwen5",
		Des:  "algorithm",
	}
	UpdateClass(cls)
}

func TestDeleteClass(t *testing.T) {
	id := 67
	DeleteClass(id)
}
