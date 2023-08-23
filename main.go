package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct{
	Title string
	Director string
}

func main() {
	fmt.Println("Hello world!")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		films := map[string][]Film{
			"Films":{
				{Title: "The God Father", Director: "Francis Form Coppola"},
				{Title: "Blade Runner", Director: "Ridely Scott"},
				{Title: "The Thing", Director: "Jhon Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")

		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)

		tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1) // Use http.HandleFunc instead of http.Handle
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
