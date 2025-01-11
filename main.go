package main

import (
	"net/http"

	"example.com/mouse-counter/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(c echo.Context, status int, cmp templ.Component) error {
	c.Response().WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return RenderComponent(c, http.StatusOK, views.Hello("World"))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
