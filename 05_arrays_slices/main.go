package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string

	// Assign values
	fruits[0] = "Apple"
	fruits[1] = "Pear"

	fmt.Println(fruits)
	fmt.Println(fruits[1])

	// declare and assign
	movies := [2]string{"Harry Potter", "Lord of the Rings"}

	fmt.Println(movies)

	// slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	// length
	fmt.Println(len(fruits))
	// range
	fmt.Println(fruits[1:2])
}
