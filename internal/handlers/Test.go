package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ Msg string }{"API LIVE!"})
}
