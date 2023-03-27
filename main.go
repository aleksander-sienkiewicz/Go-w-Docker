package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	//w, r
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	http.HandleFunc("/hi", func(response http.ResponseWriter, requeset *http.Request) {
		fmt.Fprintf(response, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8001", nil))

}
