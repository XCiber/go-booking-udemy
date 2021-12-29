package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/XCiber/go-booking-udemy/pkg/config"
	"github.com/XCiber/go-booking-udemy/pkg/handlers"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Handle("/metrics", promhttp.Handler())

	return mux
}
