package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/mazdazoomzoom/finalturnracing/pkg/config"
	"github.com/mazdazoomzoom/finalturnracing/pkg/db"
	"github.com/mazdazoomzoom/finalturnracing/pkg/models"
	"github.com/mazdazoomzoom/finalturnracing/pkg/routes"
	"github.com/mazdazoomzoom/finalturnracing/pkg/services"
)

type CustomRenderer struct {
	templates *template.Template
}

func (t *CustomRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	config.LoadEnv()

	db.Connect()
	db.Migrate(&models.Schedule{})

	services.GetSchedule()

	e := echo.New()
	render := &CustomRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = render

	routes.RegisterRoutes(e)

	port := os.Getenv("PORT")

	fmt.Println("Server is running at port", port)
	log.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
