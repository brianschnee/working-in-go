package main

import "fmt"

func main() {
	// define map map[key type]value type
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	// add kv
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)

	// delete from map
	delete(emails, "Sharon")
	fmt.Println(emails)

	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])
}
