package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	generator "placeholder-generator"
	"strconv"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tmpl, _ := template.ParseFiles("static/templates/index.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	heightStr := r.URL.Query().Get("height")
	widthStr := r.URL.Query().Get("width")

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}

	width, err := strconv.Atoi(widthStr)
	if err != nil {
		http.Error(w, "Invalid width", http.StatusBadRequest)
		return
	}

	generator.Generate(height, width, "static/images/placeholder.png")
	http.ServeFile(w, r, "static/images/placeholder.png")
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/404.html")
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	r.HandleFunc("/image", imageHandler)
	r.PathPrefix("/style/").Handler(http.StripPrefix("/style/", http.FileServer(http.Dir("static/style"))))
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Server didn't start")
		return
	}
}
