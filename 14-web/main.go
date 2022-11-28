package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home Page</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	// like a router
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting on PORT 3000")

	// like app.listen in express
	http.ListenAndServe(":3000", nil)
}
