package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginGetHandler(w, r)
	} else if r.Method == "POST" {
		loginPostHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

type Page struct {
	Email   string
	Content string
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := os.ReadFile("static/base.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("base").Parse(string(body))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	data := Page{
		Email:   "azmostofa06@gmail.com",
		Content: "breh",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {

}
