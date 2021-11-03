package main

import (
    "log"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Go to <a href='/login'>/login</a>"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method: ", r.Method)
		
		tmpl, _ := template.ParseFiles("./static/login.html")
		hi := "nothing"
		tmpl.Execute(w, hi)
	})

	
	log.Println("Listen and serve at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
