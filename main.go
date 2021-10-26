package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("static/login.html")
		hi := "nothing"
		tmpl.Execute(w, hi)
	})

	fmt.Println("We are alive on :8080")
	http.ListenAndServe(":8080", nil)
}
