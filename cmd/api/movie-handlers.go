package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(rw http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id"))
		app.errorJSON(rw, err)
		return
	}

	app.logger.Println("id is ", id)

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Println(errors.New("invalid id"))
		app.errorJSON(rw, err)
		return
	}
	err = app.writeJSON(rw, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) getAllMovies(rw http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

	err = app.writeJSON(rw, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) getAllGenres(rw http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GenresAll()
	if err != nil {
		app.errorJSON(rw, err)
		return
	}

	err = app.writeJSON(rw, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(rw, err)
		return
	}
}

func (app *application) deleteMovie(rw http.ResponseWriter, r *http.Request) {

}
func (app *application) insertMovie(rw http.ResponseWriter, r *http.Request) {

}
func (app *application) updateMovie(rw http.ResponseWriter, r *http.Request) {

}
func (app *application) searchMovies(rw http.ResponseWriter, r *http.Request) {

}
