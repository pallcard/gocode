package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!", r.URL.Path)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	//http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", mux)
	//http.ListenAndServe(":8080", nil)
}