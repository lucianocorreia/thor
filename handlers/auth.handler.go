package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/thor/views/auth"
)

func (h *Handler) HandleLogin(c echo.Context) error {
	return nil
}

func (h *Handler) HandleRegister(c echo.Context) error {
	props := auth.RegisterViewProps{
		App: getAppData(c, "Criar Conta"),
	}
	return render(c, auth.RegisterView(props))
}
