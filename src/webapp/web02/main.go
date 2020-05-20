package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandle struct {}

func (m *MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "myHandle")
}

func main() {
	myHandle := MyHandle{}
	//http.Handle("/myHandler", &myHandle)
	server := http.Server{
		Addr: ":8080",
		Handler: &myHandle,
		ReadTimeout: 2*time.Second,
	}

	server.ListenAndServe()
	//http.ListenAndServe(":8080", nil)
}