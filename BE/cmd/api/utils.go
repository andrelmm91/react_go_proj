package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type JSONresponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	//Marshalling json
	out, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshal payload into json: %s", err)
		return err
	}

	// add headers to the response if there is any
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// setting headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// writing and sending json
	_, err = w.Write(out)
	if err != nil {
		log.Printf("Error marshal payload into json: %s", err)
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		log.Printf("Error reading the json file: %s", err)
		return err
	}

	// checking if received only one json file
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		log.Printf("Error reading the json file: %s", err)
		return errors.New("body must onlz contain a  single JSON value")
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JSONresponse

	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}
