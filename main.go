package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("static/login.html")
		hi := "nothing"
		tmpl.Execute(w, hi)
	})

	dat, err := os.ReadFile("goChat/goChat/static/login.html")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("goChat/goChat/static/login.html")
	check(err)

	b1 := make([]byte, 15)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	f.Close()

	http.ListenAndServe(":8080", nil)
}
