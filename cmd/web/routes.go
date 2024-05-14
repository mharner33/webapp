package main

import (
	"net/http"

	"github.com/mharner33/webapp/pkg/config"
	"github.com/mharner33/webapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", http.HandlerFunc(handlers.Repo.Home))
	mux.HandleFunc("GET /about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
