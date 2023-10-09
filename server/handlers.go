package server

import (
	"fmt"
	"html/template"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lmoscheni/go-htmx/models"
)

func homeHandler(c echo.Context) error {
	files := []string{
		"./view/layout/base.tmpl",
		"./view/pages/home.tmpl",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	return tmpl.ExecuteTemplate(c.Response().Writer, "base", nil)
}

func examplesHandler(c echo.Context) error {
	files := []string{
		"./view/layout/base.tmpl",
		"./view/pages/examples/index.tmpl",
	}

	tmpl := template.Must(template.ParseFiles(files...))

	return tmpl.ExecuteTemplate(c.Response().Writer, "base", nil)
}

func exampleHandler(c echo.Context) error {
	name := c.Param("name")

	files := []string{
		"./view/layout/base.tmpl",
		fmt.Sprintf("./view/pages/examples/%s/index.tmpl", name),
	}

	tmpl := template.Must(template.ParseFiles(files...))

	return tmpl.ExecuteTemplate(c.Response().Writer, "base", nil)
}

func decrementCounterHandler(c echo.Context) error {
	var body models.CounterStatus
	c.Bind(&body)

	tmpl := template.Must(template.ParseFiles("./view/pages/examples/counter/htmx-fragments/count.tmpl"))

	err := body.Decrement()

	if err != nil {
		tmpl := template.Must(template.ParseFiles("./view/pages/examples/counter/htmx-fragments/decrement-error.tmpl"))
		return tmpl.Execute(c.Response().Writer, nil)
	}

	return tmpl.Execute(c.Response().Writer, body)
}

func incrementCounterHandler(c echo.Context) error {
	log.Default().Println("Incrementing count")
	var counterStatus models.CounterStatus
	c.Bind(&counterStatus)
	log.Default().Print(c.Request())

	log.Default().Print(counterStatus)

	tmpl := template.Must(template.ParseFiles("./view/pages/examples/counter/htmx-fragments/count.tmpl"))

	counterStatus.Increment()
	log.Default().Print(counterStatus)

	return tmpl.Execute(c.Response().Writer, counterStatus)
}
