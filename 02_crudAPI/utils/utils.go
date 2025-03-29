package utils

import (
	"02_crudAPI/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func GetDataFromFile(file_name string) []models.Movie {
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}

	var movies []models.Movie
	json.Unmarshal(bytes, &movies)

	return movies
}

func GetMovies(w http.ResponseWriter, r *http.Request, movies []models.Movie) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, movies *[]models.Movie) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range *movies {
		if movie.ID == params["id"] {
			*movies = append((*movies)[:index], (*movies)[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(*movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request, movies []models.Movie) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request, movies *[]models.Movie) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	*movies = append(*movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, movies *[]models.Movie) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range *movies {
		if movie.ID == params["id"] {
			*movies = append((*movies)[:index], (*movies)[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			*movies = append(*movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
