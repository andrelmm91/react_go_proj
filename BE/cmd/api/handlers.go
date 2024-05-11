package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
		fmt.Printf("Error marshal payload into json: ", err)
	}
	// settinh headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// responding
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "Highlander",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     116,
		Description: "Cool",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	rd, _ = time.Parse("2006-01-02", "1981-06-21")

	rotla := models.Movie{
		ID:          2,
		Title:       "Raiders",
		ReleaseDate: rd,
		MPAARating:  "PG13",
		RunTime:     115,
		Description: "Nice",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, rotla)

	//Marshalling json
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Sprintf("Error marshal payload into json: ", err)
	}
	// settinh headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// responding
	w.Write(out)
}
