package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mayurvpatil/go-server/pkg/config"
	"github.com/mayurvpatil/go-server/pkg/handler"
	"github.com/mayurvpatil/go-server/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
