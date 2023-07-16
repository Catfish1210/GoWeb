package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			name := r.Form.Get("name")

			data := PageData{Message: "Button Clicked: " + name}

			tmpl, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.URL.Path == "/favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			w.WriteHeader(http.StatusOK)
		} else {
			if r.URL.Path == "/styles.css" {
				http.ServeFile(w, r, "styles.css")
			} else {
				http.ServeFile(w, r, "index.html")
			}
		}
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			name := r.Form.Get("name")

			fmt.Println("Button Clicked: " + name)

			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
