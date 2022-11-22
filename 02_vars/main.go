package main

import "fmt"

func main() {
	// var name = "Brian"
	var age int32 = 37
	const isCool = true

	// shorthand
	name := "Brian"
	size := 1.3

	// multiple vars in a single line
	name, email := "brian", "brian@gmail.com"

	fmt.Println(name, email, age)
	fmt.Printf("%T\n%T\n", isCool, size)
}
