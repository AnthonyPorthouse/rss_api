package main

import (
	"app/controllers"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("resources/templates/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.Use(middleware.Logger())

	e.Static("/", "public")

	controllers.MapBaseRoutes(e)

	e.Logger.Fatal(e.Start(":1234"))
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["route"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
