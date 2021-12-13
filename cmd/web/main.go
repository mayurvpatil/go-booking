package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mayurvpatil/go-server/pkg/config"
	"github.com/mayurvpatil/go-server/pkg/handler"
	"github.com/mayurvpatil/go-server/pkg/render"
)

const portNumber = ":8080"

// main function
func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("can not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("starting server on port number " + portNumber)

	err = srv.ListenAndServe()
	log.Fatal(err)

}
