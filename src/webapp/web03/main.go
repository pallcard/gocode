package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"webapp/web03/model"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Path: ", r.URL.Path)
	fmt.Fprintln(w, "RawQuery: ", r.URL.RawQuery)
	fmt.Fprintln(w, "Header: " , r.Header)
	fmt.Fprintln(w, "Header Accept-Encoding: ", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "Header Accept-Encoding: ", r.Header.Get("Accept-Encoding"))

	//len := r.ContentLength
	//body := make([]byte, len)
	//
	//r.Body.Read(body)
	//fmt.Fprintln(w, "body: ", string(body))

	// method 1
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("err: ", err)
	}
	fmt.Fprintln(w, "form:", r.Form)
	fmt.Fprintln(w, "form:", r.PostForm)

	// method 2
	fmt.Fprintln(w, "formvalue", r.FormValue("user"))
	fmt.Fprintln(w, "formvalue", r.FormValue("username"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{
		ID: 1,
		Username: "123",
		Password: "123",
		Email: "123@163.com",
	}

	json, _ := json.Marshal(user)
	w.Write(json)
}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("index.html")
	t := template.Must(template.ParseFiles("index.html"))

	t.Execute(w, "template")

}
func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJsonRes", testJsonRes)
	http.HandleFunc("/testRedirect", testRedirect)
	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
