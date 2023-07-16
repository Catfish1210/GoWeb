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
			fmt.Println("Before Calculator Result: " + result)
			fmt.Println("After Calculator Result: ", calculateResult(result))
			w.WriteHeader(http.StatusOK)
		} else if r.URL.Path == "/favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			http.ServeFile(w, r, "favicon.ico")
		} else if r.URL.Path == "/styles.css" {
			w.Header().Set("Content-Type", "text/css")
			http.ServeFile(w, r, "styles.css")
		} else {
			http.ServeFile(w, r, "index.html")
		}
	})

	http.ListenAndServe(":8080", nil)
}

func calculateResult(input string) []string {
	var operatorIndex []int
	var values []string
	for i, char := range input {
		if char < '0' || char > '9' {
			operatorIndex = append(operatorIndex, i)
			value := input[:i]
			values = saveValue(values, value)
		}

	}
	// values := make([]int, len(operatorIndex)+1)

	return values

}

func saveValue(values []string, value string) []string {
	return append(values, value)
}
