package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("home.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":5001", r)
	if err != nil {
		panic(err.Error())
	}
}
