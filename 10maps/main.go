package main

import "fmt"

func main() {
	fmt.Println("We do MAPS")

	languages := make(map[int]int)

	languages[1] = 11
	languages[2] = 22
	languages[3] = 33
	fmt.Println("Liist of all languages: ", languages)
	fmt.Printf("Type of the DS is: %T \n", languages)
	fmt.Println(languages[1])
	for key, value := range languages {
		languages[key] = languages[key*2]
		fmt.Printf("For key %v, value is %v \n", key, value)
		fmt.Println("Liist of all languages: ", languages)
	}



}
