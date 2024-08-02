package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", greeting)
	fmt.Println("listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
