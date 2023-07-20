package main

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
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
			// calculateResult(result)
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

func calculateResult(input string) int {
	expr, err := govaluate.NewEvaluableExpression(input)
	if err != nil {
		return 0
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0
	}

	switch res := result.(type) {
	case int:
		return res
	case float64:
		return int(res)
	default:
		return 0
	}

}
