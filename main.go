package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := html.EscapeString(r.URL.Path)
		fmt.Fprintf(w, "Hello, you are %q", s)
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	fmt.Println("Server is Running")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
