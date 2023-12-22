package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-fullstack/views"
)

func (h *Handler) HandleIndex(c echo.Context) error {
	props := views.IndexViewProps{
		App: getAppData(c, "Index"),
	}
	return render(c, views.IndexView(props))
}
