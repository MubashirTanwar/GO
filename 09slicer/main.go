package main

import (
	"fmt"
)

func main() {
	fmt.Println("We are learning slicers")

	var fruits = []string{"Apple", "bannaa", "PEar"}

	fruits = append(fruits[:], "Mango", "Strawaberry", fruits[1])
	fruits[0] = ""
	fmt.Printf("type of Fruits is %T \n", fruits)
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 23
	highScores[1] = 23458304
	highScores[2] = 298786
	highScores[3] = 24

	fmt.Println(len(highScores))
	// highScores = append(highScores, 83764, 345, 234325)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// fmt.Println(highScores)

	var courses = []string{"recat", "nextjs", "golund"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)




}
