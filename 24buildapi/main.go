package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	return c.Name == ""
}

func main() {
	fmt.Println("WE ARE MAKING API")
	r:= mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")

	courses = append(courses , Courses{
		Id: "1", Name: "React", Price: 1000, Author: &Author{Name: "Mubashir",Email: "tanwar0210@gmail.com"},
	})

	courses = append(courses , Courses{
		Id: "2", Name: "NodeJS", Price: 500, Author: &Author{Name: "Rahul",Email: "rahul0210@gmail.com"},
	})

	r.HandleFunc("/courses", getAllCourses).Methods("GET")

	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")

	r.HandleFunc("/course", createOneCourse).Methods("POST")
	
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":4000", r))
	
}

// controllers ->

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the home route</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	params := mux.Vars(r)
	fmt.Println(params)

	for _, v := range courses {
		if v.Id == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We are in get one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("No data in body")
		return
	}
	var course Courses
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("It is empty")
		return
	}

	rand.New(rand.NewSource(13))

	// StringWithChar
	// new(NewSource(time.Now().UnixNano()))
	course.Id = strconv.Itoa(rand.Intn(10))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}




func updateCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("We are in get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if len(params) == 0 {
		json.NewEncoder(w).Encode("Add Id in params")
	}

	if params["id"] == "" {
		json.NewEncoder(w).Encode("Add Id in params")
	}

	if r.Body == nil {
		json.NewEncoder(w).Encode("Add body")
	}

	var course Courses 
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Add name")
	}

	// get idx of the course
	var id string = params["id"]
	var target int
	for i,v := range courses {
		if v.Id == id {
			target = i
			break
		}
	}

	// update course via idx
	courses[target] = course

}


func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("We are in delete course")

	var params = mux.Vars(r)
	if params["id"] == ""{
		json.NewEncoder(w).Encode("Add ID in params")
	}

	var id string = params["id"]
	var target int
	for i,v := range courses {
		if v.Id == id {
			target = i
			break
		}
	}

	// TODO: handle target not found

	c := append(courses[:target], courses[target+1:]...)
	
	json.NewEncoder(w).Encode(c)
	
}

