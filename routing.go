package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/student", CreateStudent).Methods("POST")
	r.HandleFunc("/students", GetStudents).Methods("GET")
	r.HandleFunc("/student/{id}", GetStudentByID).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", r))
}
