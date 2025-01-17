package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/about", AboutUs)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
