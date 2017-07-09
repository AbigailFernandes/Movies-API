package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	rating, _ := strconv.Atoi(r.URL.Query().Get("ratinggeq"))
	movies := getMovies(rating)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		panic(err)
	}
}

func MovieIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID, _ := strconv.Atoi(vars["movieID"])
	movie := getMoviesById(movieID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		panic(err)
	}
}
