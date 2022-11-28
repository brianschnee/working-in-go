package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	age       int

	// firstName, lastName, city string
	// age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// changeLastName (pointer reciever)
func (p *Person) changeLastName(changeLastName string) {
	p.lastName = changeLastName
}

func main() {
	// Init person using struct
	john := Person{firstName: "John", lastName: "Doe", city: "Boston", age: 34}

	// // alternate
	// jane := Person{"Jane", "Doe", "Boston", 34}

	// fmt.Println(john, jane)
	// john.age++
	// fmt.Println(john.age)

	john.hasBirthday()
	john.changeLastName("Hello")
	fmt.Println(john.greet())
}
