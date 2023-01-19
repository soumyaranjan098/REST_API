package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studs []Student
	Database.Find(&studs)
	json.NewEncoder(w).Encode(studs)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	Database.First(&student, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stud Student
	json.NewDecoder(r.Body).Decode(&stud)
	Database.Create(&stud)
	json.NewEncoder(w).Encode(stud)
}
