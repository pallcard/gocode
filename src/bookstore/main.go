package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/",http.FileServer(http.Dir("view/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/",http.FileServer(http.Dir("view/pages"))))
	http.HandleFunc("/main", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
