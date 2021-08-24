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
	app.writeJSON(rw, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(rw http.ResponseWriter, r *http.Request) {

}
