package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello,world!")
	return c.HTML(http.StatusOK, "Hello,world!")
}
