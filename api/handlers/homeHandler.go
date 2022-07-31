package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func Home(c echo.Context) error {
	data := M{"message": "Hello World!"}
	return c.Render(http.StatusOK, "index.html", data)
}
