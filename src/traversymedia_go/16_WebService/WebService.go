package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h3>Hello World , via GO Web !!!!</h3>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page ")
}
func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Printf("Starting webservic ........")
	http.ListenAndServe(":3000", nil)
}
