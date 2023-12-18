package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func mainHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/index.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/404.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.PathPrefix("/style/").Handler(http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Server didn't start")
		return
	}
}
