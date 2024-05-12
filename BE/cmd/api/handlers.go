package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// creating payload and populating it
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Up and running",
		Version: "1.0.0",
	}

	//Marshalling json
	out, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshal payload into json: %s", err)
	}
	// settinh headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// responding
	w.Write(out)
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	log.Println(movies)
	if err != nil {
		log.Printf("Error loading AllMovies function: %s", err)
		return
	}

	//Marshalling json
	out, err := json.Marshal(movies)
	if err != nil {
		log.Printf("Error marshal payload into json: %s", err)
	}
	// settinh headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// responding
	w.Write(out)
}
