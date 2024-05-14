package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mharner33/webapp/pkg/config"
	"github.com/mharner33/webapp/pkg/handlers"
	"github.com/mharner33/webapp/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache!")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Printf("Starting webserver on port: %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to start webserver")
	}
}
