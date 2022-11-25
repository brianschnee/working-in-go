package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 21, 4}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("Index: %d - Id: %d\n", i, id)
	}

	// not using index (use _ if you arent using the value, _ replaces whatever variable would hold the index in this case)
	for _, id := range ids {
		fmt.Printf("Id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	// keys and values
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// keys
	for k := range emails {
		fmt.Printf("name: " + k)
	}
}
