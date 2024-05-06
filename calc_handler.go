package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ErrorMessage struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type OperationError struct {
	Result    string `json:"result"`
	Operation string `json:"op"`
}

type OperationResult struct {
	Result    float64 `json:"result"`
	Operation string  `json:"op"`
}

var operators = map[string]Calcular{
	"*":   &Multiplicacao{},
	"mul": &Multiplicacao{},
	"-":   &Subtracao{},
	"sub": &Subtracao{},
	"+":   &Soma{},
	"sum": &Soma{},
	"/":   &Divisao{},
	"div": &Divisao{},
	"^":   &Potencia{},
	"pow": &Potencia{},
	"&":   &Radiciacao{},
	"rot": &Radiciacao{},
}

type calc_Handler struct{}

func (calc_Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	log.Print("[INFO]" + req.Method + " " + req.URL.Path)

	if req.Method != "GET" || req.URL.Path != "/result" {
		e := ErrorMessage{Code: 404, Error: "Not Found"}
		r, _ := json.Marshal(e)
		resp.WriteHeader(404)
		resp.Write(r)
		return
	}

	operation := req.URL.Query().Get("op")

	if !isOperationValid(operation) {
		c := OperationError{Result: "invalid expression", Operation: operation}
		r, _ := json.Marshal(c)
		resp.WriteHeader(400)
		resp.Write(r)
		return
	}

	operator := getOperatorFromString(operation)
	numbers := strings.Split(operation, operator)

	mathe := operators[operator]
	x, _ := strconv.ParseFloat(numbers[0], 64)
	y, _ := strconv.ParseFloat(numbers[1], 64)

	result := mathe.Calcular(x, y)

	c := OperationResult{Result: result, Operation: operation}
	r, _ := json.Marshal(c)
	resp.WriteHeader(200)
	resp.Write(r)

}

func isOperationValid(op string) bool {
	for k := range operators {
		if strings.Contains(op, k) {
			return true
		}
	}

	return false
}

func getOperatorFromString(op string) string {
	for k := range operators {
		if strings.Contains(op, k) {
			return k
		}
	}

	return ""
}
