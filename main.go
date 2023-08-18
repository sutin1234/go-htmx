package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello Go HTMX")

	http.HandleFunc("/", handleFunc1)
	http.HandleFunc("/add-film/", handleFunc2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleFunc1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "Golang", Director: "Thinny"},
			{Title: "Rust", Director: "Thin"},
			{Title: "Angular", Director: "Sutin"},
		},
	}
	tmpl.Execute(w, films)
}
func handleFunc2(w http.ResponseWriter, r *http.Request) {
	log.Println("HTMX Request")

	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	htmlStr := fmt.Sprintf("<li>%s -  %s</li>", title, director)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}
