package main

import (
	"fmt"
	"html/template"
	"time"

	// "io"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			/*
				variable: films, of type "map",
				keys are type "String" and values are the type "list"
			*/
			"Films": {
				{Title: "A Clockwork Orange", Director: "Stanley Kubric"},
				{Title: "The Godfather Trilogy", Director: "Francis Ford Coppola"},
				{Title: "Interstellar", Director: "Christopher Nolan"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director}) // refrences to "film-list-element" block in index.html

		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		// tmpl.Execute(w, nil)

		// log.Print("HTMX Request Received")
		// log.Print(r.Header.Get("HX-Request"))
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)
	log.Fatal(http.ListenAndServe(":8010", nil))

}
