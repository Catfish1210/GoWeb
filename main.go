package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		fmt.Println("Submitted name:", name)
		w.Write([]byte("Hello, " + name + "!"))
	})

	http.ListenAndServe(":8080", nil)
}
