package main

import "fmt"

func main() {
	fmt.Println("We are learning struct")

	mubashir := User{"Mubashir","trash@test", true, 20,  Address{"104", "Mumbai", true, 401107} }

	fmt.Printf("User details are: %v %v \n", mubashir.Name, mubashir.Age)
	fmt.Printf("User details are: %+v \n", mubashir)
	mubashir.GetAddress()
}


type User struct {
	Name string
	Email string
	Status bool
	Age int
	Address Address
}

type Address struct {
	Room string
	City string
	Rented bool
	Pincode int
}

func (u User) GetAddress(){
	fmt.Printf("Address is: %v\n City is: %v\n", u.Address, u.Address.City )
}