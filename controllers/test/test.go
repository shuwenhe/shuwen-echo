package test

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index2(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "test123...")
}
