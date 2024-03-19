package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     string `json:"year"`
}

var movies = []Movie{
	{ID: "1", Title: "The Shawshank Redemption", Director: "Frank Darabont", Year: "1994"},
	{ID: "2", Title: "The Godfather", Director: "Francis Ford Coppola", Year: "1972"},
	{ID: "3", Title: "The Dark Knight", Director: "Christopher Nolan", Year: "2008"},
}

func getALlMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request) {

	// url is localhost:8081/movie/1
	// r.URL.Path will be /movie/1
	// strings.TrimPerfix will remove the r.URL.Path and rest will be assigned to id
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/movie/")

	for i := 0; i < len(movies); i++ {
		if movies[i].ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.NotFound(w, r)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/delete/")

	for i := 0; i < len(movies); i++ {
		if movies[i].ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie delete successfully"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"Error": "Movie not found"})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	w.Header().Set("Content-Type", "application.json")
	var newmovie Movie

	if err := json.NewDecoder(r.Body).Decode(&newmovie); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "error in decoding"})
		return
	}
	movies = append(movies, newmovie)
	json.NewEncoder(w).Encode(map[string]string{"success": "movie update successfully"})

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/update/")
	w.Header().Set("Content-Type", "application/json")
	for i := 0; i < len(movies); i++ {
		if movies[i].ID == id {
			// delete the current id
			movies = append(movies[:i], movies[i+1:]...)

			// add the new one
			var newmovie Movie
			if err := json.NewDecoder(r.Body).Decode(&newmovie); err != nil {
				json.NewEncoder(w).Encode(map[string]string{"error": "error in decoding"})
				return
			}
			movies = append(movies, newmovie)
			json.NewEncoder(w).Encode(map[string]string{"success": "movie updated successfully"})
			return
		}
	}
	json.NewEncoder(w).Encode(map[string]string{"error": "movie not found"})
}

func main() {
	http.HandleFunc("/movies", getALlMovies)
	http.HandleFunc("/movie/", getMovieById)
	http.HandleFunc("/delete/", deleteMovie)
	http.HandleFunc("/create", createMovie)
	http.HandleFunc("/update/", updateMovie)
	fmt.Println("Server starting at port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
