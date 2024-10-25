package main

import (
	"25mongodb/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("We are connecting mongo db to go")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening to port 4000")
}
