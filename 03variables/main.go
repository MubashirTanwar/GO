package main

import "fmt"

const LoginToken string = "gibberish" //public
const logouttoken string = "lol" //private

func main() {
	var username string = "Mubashir"
	fmt.Println("Variable is of type:  \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallNum int = 255
	fmt.Printf("Variable is of type: %T \n", smallNum)

	// Default
	var initString string
	fmt.Println(initString)
	fmt.Printf("Variable is of type: %T \n", initString)

	// implicit type
	var website = "mubashir.com"
	fmt.Println(website)

	// no vars style
	// not allowed outside methods
	
	numOfUsers := "2354.864876"
	fmt.Println(numOfUsers)
	fmt.Printf("Variable is of type: %T \n", numOfUsers)

	fmt.Println(logouttoken)
	fmt.Printf("Variable is of type: %T \n", logouttoken)

}
