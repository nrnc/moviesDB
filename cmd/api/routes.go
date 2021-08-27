package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc("GET", "/status", app.statusHandler)
	router.HandlerFunc("GET", "/v1/movies/:id", app.getOneMovie)
	router.HandlerFunc("GET", "/v1/movies", app.getAllMovies)
	router.HandlerFunc("GET", "/v1/genres", app.getAllGenres)
	router.HandlerFunc("GET", "/v1/genres/:id", app.getAllMoviesByGenre)
	return app.enableCORS(router)
}
