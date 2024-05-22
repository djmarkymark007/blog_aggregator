package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type errRespond struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	ret := errRespond{Error: msg}
	data, err := json.Marshal(ret)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(data))
}

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func readiness(w http.ResponseWriter, r *http.Request) {
	type resType struct {
		Status string `json:"status"`
	}
	res := resType{Status: "ok"}
	respondWithJson(w, 200, res)
}

func errTest(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	serverHandler := http.NewServeMux()
	serverHandler.HandleFunc("GET /v1/readiness", readiness)
	serverHandler.HandleFunc("GET /v1/err", errTest)

	server := http.Server{Handler: serverHandler, Addr: ":" + port}
	fmt.Printf("Starting server on port: %s\n", port)
	server.ListenAndServe()
}
