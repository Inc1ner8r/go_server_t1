package main

import (
	"fmt"
	"net/http"
)

// handler function
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	//register http handler function to root url
	http.HandleFunc("/", hello)

	//serve on port 8080
	http.ListenAndServe(":8080", nil)
}
