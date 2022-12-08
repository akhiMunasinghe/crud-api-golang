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
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

// get all the available movies
func getMovies(respsonse http.ResponseWriter, request *http.Request) {
	respsonse.Header().Set("Content-Type", "application/json")
	json.NewEncoder(respsonse).Encode(movies)
}

// delete the movie matching the movie id
func deleteMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movieID := mux.Vars(request)["id"]
	//loop through the movies array until the id is found and remove the movie by combining movies before the id and after id
	for index, item := range movies {
		if item.Id == movieID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(response).Encode(movies)
}

// get the specified movie from the array
func getMovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	movieID := mux.Vars(request)["id"]
	for _, movie := range movies {
		if movie.Id == movieID {
			json.NewEncoder(response).Encode(movie)
			break
		}
	}
}

// create a new movie and add to the array
func createMovie(resposnse http.ResponseWriter, request *http.Request) {
	resposnse.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(200))
	movies = append(movies, movie)
	json.NewEncoder(resposnse).Encode(movies)
}

// update the movie in array that match the given ID
func updateMovie(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-Type", "applicayion/json")
	// movieID := mux.Vars(request)["id"]
	// for _, item := range movies {
	// 	if item.Id == movieID {
	// 		fmt.Println(request.Body)
	// 		json.NewDecoder(request.Body).Decode(&item)
	// 		break
	// 	}
	// }
	// json.NewEncoder(response).Encode(movies)
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(response).Encode(movies)
		}
	}
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "12345", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "56789", Title: "Movie Two", Director: &Director{FirstName: "Sam", LastName: "Smith"}})

	router.HandleFunc("/movies", getMovies).Methods("GET")           //get all movies in array
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")       //get the movie with correct ID
	router.HandleFunc("/movies", createMovie).Methods("POST")        //create a new movie
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    //update the movie with given ID
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") //delete the movie with the given ID

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))

}
