package main

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Branch string `json:"branch"`
	Email  string `json:"email"`
}
