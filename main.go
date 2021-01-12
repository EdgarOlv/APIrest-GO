package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes
	
	s.HandleFunc("/movies", GetAll).Methods("GET")
	s.HandleFunc("/movies/{id}", GetByID).Methods("GET")
	s.HandleFunc("/movies/", Create).Methods("POST")
	s.HandleFunc("/movies/{id}", Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}
