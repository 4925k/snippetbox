package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// set up middleware
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// init a new server mux and handle the routes
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	// mux.HandleFunc("/snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)

	// // create a file server
	// fileServer := http.FileServer(http.Dir("./ui/static"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir(".ui/static/"))
	mux.Get("/static", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
