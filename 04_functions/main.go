package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// func sum(a int, b int) int {
// 	return a + b
// }

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Brian"))
	fmt.Println(sum(3, 4))
}
