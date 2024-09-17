package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// DATA MODELLING
type Courses struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var courses []Courses

// MIDDLE WARE
func (c *Courses) isEmpty() bool {
	return c.Id == "" && c.Name == ""
}

func main() {
	fmt.Println("WE ARE MAKING API")
}

// controllers ->

func home(w http.ResponseWriter, r http.Request) {
	w.Write([]byte("<h1>This is the home route</h1>"))
}

func getAllCourses(w http.ResponseWriter, r http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course")
	params := mux.Vars(r)
	fmt.Println(params)

	for _ ,v := range courses{
		if v.Id == params["id"]{
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}