package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucianocorreia/go-fullstack/handlers"
)

func main() {
	app := echo.New()

	h := handlers.Handler{}

	app.GET("/", h.HandleIndex)
	app.Static("/static", "static")

	app.Logger.Info("Starting server on port 3000")
	app.Start(":3000")
}
