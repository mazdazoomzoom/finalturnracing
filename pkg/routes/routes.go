package routes

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/", home)
}

func home(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}
