package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello World</h1>")


}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on http://localhost:4000/")
	http.ListenAndServe(":4000", nil)
}