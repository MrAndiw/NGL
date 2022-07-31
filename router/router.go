package router

import (
	"NGL/api/handlers"
	"NGL/api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {

	e := echo.New()

	// create group
	g := e.Group("/api/v1")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddleware(g)

	// set main router
	MainRouter(e)

	// set group router
	AdminRouter(g)

	return e
}

func MainRouter(e *echo.Echo) {
	e.POST("/SetQuestion", handlers.SetQuestion)
	e.POST("/GetQuestion", handlers.GetQuestion)

	e.GET("/login", handlers.Login)
}

func AdminRouter(g *echo.Group) {
	g.GET("/dashboard", handlers.GetDashboard)
}
