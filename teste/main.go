package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/ninja":
		fmt.Fprint(w, "GoLang Dojo")
	default:
		fmt.Fprint(w, "Error!")
	}

}

func helloWorldHtmlPage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func timeOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accesando...")
	time.Sleep(2 * time.Second)
}

func main() {
	http.HandleFunc("/", helloWorldHtmlPage)
	http.HandleFunc("/timeout", timeOut)
	http.ListenAndServe(":8080", nil)

	// server := http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      nil,
	// 	ReadTimeout:  1000,
	// 	WriteTimeout: 1000,
	// }
	// server.ListenAndServe()
}
