package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("start server")
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello</h1>")
}
