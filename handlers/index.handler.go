package handlers

import (
	"github.com/labstack/echo/v4"
	auth "github.com/lucianocorreia/go-fullstack/views"
)

// HandleIndex handles the index route
func (h *Handler) HandleIndex(c echo.Context) error {
	props := auth.IndexViewProps{
		App: getAppData(c, "Index"),
	}

	return render(c, auth.IndexView(props))
}
