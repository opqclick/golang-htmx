package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct{
	Title string
	Director string
}

func main() {
	fmt.Println("Hello world!")

	myHandler := func(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/", myHandler) // Use http.HandleFunc instead of http.Handle

	log.Fatal(http.ListenAndServe(":8000", nil))
}
