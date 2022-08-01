package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func Home(c echo.Context) error {
	data := M{
		"title":   "Welcome To NGL",
		"message": "Please Share Your link to Instagram",
	}
	return c.Render(http.StatusOK, "index.html", data)
}
