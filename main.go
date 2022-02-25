package main

import (
	"fmt"
	"log"
	"net/http"

	S "./services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

    // Init the mux router

    router := mux.NewRouter()

	// Route handles & endpoints

    // Get all movies
 
    router.HandleFunc("/movies/", S.GetMovies).Methods("GET")

    // Create a movie
    router.HandleFunc("/movies/", S.CreateMovie).Methods("POST")

    // Delete a specific movie by the movieID
    router.HandleFunc("/movies/{movieid}", S.DeleteMovie).Methods("DELETE")


    // serve the app
    fmt.Println("Server at 8080")

    log.Fatal(http.ListenAndServe(":8080", router))
}



