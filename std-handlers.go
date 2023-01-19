package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studs []Student
	res := Database.Find(&studs)
	if res.RowsAffected != 0 {
		json.NewEncoder(w).Encode(studs)
	} else {
		json.NewEncoder(w).Encode("No student found..")
	}

}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	res := Database.First(&student, mux.Vars(r)["id"])
	if res.RowsAffected != 0 {
		json.NewEncoder(w).Encode(student)
	} else {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("No student found with the given id")

	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stud Student
	json.NewDecoder(r.Body).Decode(&stud)
	Database.Create(&stud)
	json.NewEncoder(w).Encode("New Student created Successfully...")

}
