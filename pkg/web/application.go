package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

type Application struct {
	Router *chi.Mux
}

func NewWebApplication() *Application {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return &Application{Router: r}
}

func (a *Application) Run() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", a.Router)

}
