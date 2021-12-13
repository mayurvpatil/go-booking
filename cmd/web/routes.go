package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mayurvpatil/go-server/pkg/config"
	"github.com/mayurvpatil/go-server/pkg/handler"
)

func routes(app *config.AppConfig) http.Handler {

	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handler.Repo.About))

	return mux

}
