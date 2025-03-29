package main

import (
	"02_crudAPI/models"
	"02_crudAPI/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies []models.Movie

func main() {

	movies := utils.GetDataFromFile("data.json")

	r := mux.NewRouter()

	r.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		utils.GetMovies(w, r, movies)
	}).Methods("GET")

	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		utils.GetMovie(w, r, movies)
	}).Methods("GET")
	r.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		utils.CreateMovie(w, r, &movies)
	}).Methods("POST")
	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		utils.UpdateMovie(w, r, &movies)
	}).Methods("PUT")
	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		utils.DeleteMovie(w, r, &movies)
	}).Methods("DELETE")

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
