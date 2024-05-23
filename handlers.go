package main

import (
	"net/http"
)

func readiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, map[string]string{"status": "ok"})
}

func errTest(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
