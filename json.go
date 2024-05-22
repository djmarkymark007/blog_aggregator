package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errRespond struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errRespond{Error: msg})
}

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Add("content-type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}
