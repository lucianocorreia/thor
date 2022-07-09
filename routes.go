package thor

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (t *Thor) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if t.IsDebug {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer)

	return mux
}
