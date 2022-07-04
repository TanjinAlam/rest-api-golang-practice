package main

import (
	"strconv"
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
)


type Movie struct{
	ID string `json: "id"`
	ISBN string `json: "id"`
	Title string `json: "id"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:] ...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
	return 
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
	return 
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:] ...)
			var movie Movie	
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["ID"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
	return 
}

func main() {
	//addding moc data
	movies = append(
		movies, Movie{ID: "1", ISBN:"438227", Title: "Movie One", 
		Director : &Director{Firstname: "Tanjin", Lastname:"Alam"}})
	
	r := mux.NewRouter()
	//http://localhost:8000/movies
    r.HandleFunc("/movies", getMovies).Methods("GET")
	//http://localhost:8000/movies/1
    r.HandleFunc("/movies/{ID}", getMovie).Methods("GET")
	//http://localhost:8000/create-movies
    r.HandleFunc("/create-movies", createMovie).Methods("POST")
	//http://localhost:8000/update-movies/1
	r.HandleFunc("/update-movies/{ID}", updateMovie).Methods("PUT")
	//http://localhost:8000/delete-movies/1
    r.HandleFunc("/delete-movies/{ID}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
    http.Handle("/", r,)
}
