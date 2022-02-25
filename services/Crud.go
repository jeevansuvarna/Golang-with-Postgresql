package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	J "../JsonResponse"
	D "../db"
	H "../helper"
	L "../structs"
	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
 
    db := D.SetupDB()

    H.PrintMessage("Getting movies...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM movies")

    // check errors
    H.CheckErr(err)

    // var response []JsonResponse
    var movies []L.Movie

    // Foreach movie
    for rows.Next() {
        var id int
        var movieID string
        var movieName string

        err = rows.Scan(&id, &movieID, &movieName)

        // check errors
        H.CheckErr(err)

        movies = append(movies, L.Movie{MovieID: movieID, MovieName: movieName})
    }

    var response = J.JsonResponse{Type: "success", Data: movies}

    json.NewEncoder(w).Encode(response)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
    movieID := r.FormValue("movieid")
    movieName := r.FormValue("moviename")

    var response = J.JsonResponse{}

    if movieID == "" || movieName == "" {
        response = J.JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
    } else {
        db := D.SetupDB()

        H.PrintMessage("Inserting movie into DB")

        fmt.Println("Inserting new movie with ID: " + movieID + " and name: " + movieName)

        var lastInsertID int
    err := db.QueryRow("INSERT INTO movies(movieID, movieName) VALUES($1, $2) returning id;", movieID, movieName).Scan(&lastInsertID)

    // check errors
    H.CheckErr(err)

    response = J.JsonResponse{Type: "success", Message: "The movie has been inserted successfully!"}
    }

    json.NewEncoder(w).Encode(response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    movieID := params["movieid"]

    var response = J.JsonResponse{}

    if movieID == "" {
        response = J.JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
    } else {
        db := D.SetupDB()

        H.PrintMessage("Deleting movie from DB")

        _, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)

        // check errors
        H.CheckErr(err)

        response = J.JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
    }

    json.NewEncoder(w).Encode(response)
}

