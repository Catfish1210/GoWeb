package main

import (
	"fmt"
	"net/http"
)

type PageData struct {
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Handle the POST request from JavaScript (getResult function)
			r.ParseForm()
			result := r.Form.Get("result")

			fmt.Println("Calculator Result: " + result)

			// Respond with a success status code (200 OK) to the AJAX request.
			w.WriteHeader(http.StatusOK)
		} else if r.URL.Path == "/favicon.ico" {
			// Handle the request for favicon.ico
			w.Header().Set("Content-Type", "image/x-icon")
			http.ServeFile(w, r, "favicon.ico")
		} else if r.URL.Path == "/styles.css" {
			// Handle the request for styles.css
			w.Header().Set("Content-Type", "text/css")
			http.ServeFile(w, r, "styles.css")
		} else {
			// Serve the initial HTML page for all other requests
			http.ServeFile(w, r, "index.html")
		}
	})

	http.ListenAndServe(":8080", nil)
}
