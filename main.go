package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	//"net/url"
	// "strings"
	//"bytes"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Hello. Go to <a href='/login'>/login</a>"))
		w.Write([]byte("Hello. Go to /login"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method: ", r.Method)

		r.ParseForm()
		for key, value := range r.Form {
			log.Printf("%s = %s", key, value)
		}

		// endpoint := "http://localhost:8080/login"
		// data := url.Values{r.Method}
		// r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(data)

		// buffer := new(bytes.Buffer)
		// params := url.Values{}
		// params.Set("username", "a")
		// params.Set("password", "b")
		// buffer.WriteString(params.Encode())

		tmpl, _ := template.ParseFiles("./static/login.html")
		hi := "nothing"
		tmpl.Execute(w, hi)
	})

	log.Println("Listen and serve at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
