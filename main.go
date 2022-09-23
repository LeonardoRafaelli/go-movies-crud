package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode("Movie not found!")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	// movie := json.NewDecoder(r.Body).Decode(&movies)
	var movie Movie

	for index, item := range movies {
		if item.ID == params["id"] {

			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(req.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(res).Encode(movie)
			return
		}
	}

}

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie Good", Director: &Director{Firstname: "Leonardo", Lastname: "Rafaelli"}})
	movies = append(movies, Movie{ID: "2", Isbn: "758990", Title: "Movie Bad", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	routes := mux.NewRouter()

	routes.HandleFunc("/movies", getMovies).Methods("GET")
	routes.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	routes.HandleFunc("/movies", createMovie).Methods("POST")
	routes.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	routes.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}

// custos (hrs) (500hrs desenv)
// qts pessoas, maquinas, servidores
// prazo, tempo de des

//custo
//tempo
//recursos
