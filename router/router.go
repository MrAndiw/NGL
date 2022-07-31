package router

import (
	"NGL/api/handlers"
	"NGL/api/middlewares"

	"github.com/labstack/echo"
)

func Start() *echo.Echo {

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
	e.POST("/CreateQuestion", handlers.CreateQuestion)
	e.POST("/GetQuestion", handlers.GetQuestion)
	e.POST("/DeleteQuestion", handlers.DeleteQuestion)
	e.POST("/ReadQuestion", handlers.ReadQuestion)

	e.GET("/login", handlers.Login)
}

func AdminRouter(g *echo.Group) {
	g.GET("/dashboard", handlers.GetDashboard)
}
