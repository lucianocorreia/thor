package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render is a helper function to render a component.
func render(c echo.Context, cmp templ.Component) error {
	return cmp.Render(c.Request().Context(), c.Response())
}
