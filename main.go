package main

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode"
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
			calculateResult(result)
			// fmt.Println("After Calculator Result: ", calculateResult(result))
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

func calculateResult(input string) {
	fmt.Println(Calculate(input))
	// var num string
	// var numSlice []string
	// operatorMap := map[string][]int{
	// 	"+": {},
	// 	"-": {},
	// 	"÷": {},
	// 	"%": {},
	// 	"x": {},
	// }

	// for i, char := range input {
	// 	fmt.Println("Iteration: ", i, "On char: ", string(char))

	// 	if isOperator(char) {
	// 		fmt.Println("-isOperator: True", string(char))
	// 		switch string(char) {
	// 		case "+":
	// 			operatorMap["+"] = append(operatorMap["+"], i)
	// 		case "-":
	// 			operatorMap["-"] = append(operatorMap["-"], i)
	// 		case "÷":
	// 			operatorMap["÷"] = append(operatorMap["÷"], i)
	// 		case "%":
	// 			operatorMap["%"] = append(operatorMap["%"], i)
	// 		case "x":
	// 			operatorMap["x"] = append(operatorMap["x"], i)
	// 		default:
	// 			fmt.Println("Invalid operator")
	// 		}

	// 		numSlice = append(numSlice, num)
	// 		num = ""
	// 	} else {
	// 		num += string(char)
	// 		if i == len(input)-1 {
	// 			numSlice = append(numSlice, num)
	// 		}
	// 	}
	// 	fmt.Println("--Num string at end of for loop: ", num)
	// 	fmt.Println("--NumSlice string at end of for loop: ", numSlice)
	// }

	// fmt.Println("Num slice: ", numSlice)
	// fmt.Println("Operator map: ", operatorMap)

	// Calculate(AtoiSlice(numSlice), operatorMap)

}

func Calculate(expression string) int {
	var num int
	var operator = '+'
	var result int
	var total int

	for _, char := range expression {
		if unicode.IsDigit(char) {
			digit, _ := strconv.Atoi(string(char))
			num = num*10 + digit
		} else if isOperator(char) {

			result = applyOperation(result, operator, num)
			total += result
			num = 0
			operator = char
		} else {
			return 0
		}
	}
	result = applyOperation(result, operator, num)

	total += result
	return total
}

func applyOperation(currentResult int, operator rune, num int) int {
	switch operator {
	case '+':
		return currentResult + num
	case '-':
		return currentResult - num
	case 'x':
		return currentResult * num
	case '÷':
		return currentResult / num
	default:
		return num
	}
}

func isOperator(char rune) bool {
	opCheck := string(char)
	if opCheck == "+" || opCheck == "-" || opCheck == "÷" || opCheck == "%" || opCheck == "x" {
		return true
	} else {
		return false
	}
}

func AtoiSlice(numSlice []string) []int {
	var sliceInt []int

	for _, number := range numSlice {
		res := 0
		for _, ch := range number {
			digit := int(ch - '0')
			res = res*10 + digit
		}
		sliceInt = append(sliceInt, res)
	}

	return sliceInt
}
