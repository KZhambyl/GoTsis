package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"GoTsis/pkg"
)

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/generals", handler.GetGenerals).Methods("GET")
	rtr.HandleFunc("/generals/{name}", handler.GetGeneralByName).Methods("GET")
	rtr.HandleFunc("/health", handler.HealthCheck).Methods("GET")

	log.Println("Server is starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", rtr))
}