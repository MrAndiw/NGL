package router

import (
	"NGL/api/handlers"
	"NGL/api/middlewares"
	"io"
	"text/template"

	"github.com/labstack/echo"
)

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

// == ROUTER ==
func Start() *echo.Echo {

	e := echo.New()

	e.Renderer = NewRenderer("template/*.html", true)

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
	e.GET("/", handlers.Home)
	e.POST("/CreateQuestion", handlers.CreateQuestion)
	e.POST("/GetQuestion", handlers.GetQuestion)
	e.POST("/DeleteQuestion", handlers.DeleteQuestion)
	e.POST("/ReadQuestion", handlers.ReadQuestion)

	e.GET("/login", handlers.Login)
}

func AdminRouter(g *echo.Group) {
	g.GET("/dashboard", handlers.GetDashboard)
}
