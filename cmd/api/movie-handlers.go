package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/nchukkaio/moviesDB/models"
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

	movie := models.Movie{
		ID:          id,
		Title:       "Some Movie",
		Description: "Some Description",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		UpdatedAt:   time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
	}
	app.writeJSON(rw, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(rw http.ResponseWriter, r *http.Request) {

}
