package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-fullstack/models"
)

// render is a helper function to render a component.
func render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}

// getAppData is a helper function to get the app data.
func getAppData(c echo.Context, title string) models.App {
	return models.App{
		Title:           title,
		IsAuthenticated: false,
		RoutePath:       c.Path(),
	}

}
