package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-fullstack/views"
)

func (h *Handler) HandleIndex(c echo.Context) error {
	props := views.IndexViewProps{
		Title: "Painel de Controle",
	}
	return render(c, views.IndexView(props))
}
