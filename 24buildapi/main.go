package main

import "fmt"


// DATA MODELLING 
type Courses struct {
	Id 		string  `json:"id"`
	Name 	string  `json:"name"`
	Price 	int `json:"price"`
	Author *Author `json:"author"`

}

type Author struct {
	Name 	string `json:"name"`
	Email 	string `json:"email"`
}


// MIDDLE WARE
func (c Courses) isEmpty() bool {
	return c.Id == "" && c.Name == ""
}

func main() {
	fmt.Println("WE ARE MAKING API")
}


