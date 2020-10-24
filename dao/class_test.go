package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-echo/models"
)

func testAddClass(t *testing.T) {
	cls := &models.Class{
		Name: "shuwen",
		Des:  "mathematics",
	}

	cls2 := &models.Class{
		Name: "shuwen",
		Des:  "algorithm",
	}
	AddClass(cls)
	AddClass(cls2)
}

func testUpdateClass(t *testing.T) {
	cls := &models.Class{
		ID:   5,
		Name: "Richard",
		Des:  "son",
	}
	UpdateClass(cls)
}

func testGetClasses(t *testing.T) {
	classes, _ := GetClasses()
	for _, class := range classes {
		fmt.Println("class = ", class)
	}
}
