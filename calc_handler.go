package main

import (
	"encoding/json"
	"errors"
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
	Result    string `json:"result"`
	Operation string `json:"op"`
}

var operators = map[string]Mathematician{
	"*":   Mult{},
	"mul": Mult{},
	"+":   Sum{},
	"sum": Sum{},
	"-":   Sub{},
	"sub": Sub{},
	"/":   Div{},
	"div": Div{},
	"^":   Pow{},
	"pow": Pow{},
	"&":   Rot{},
	"rot": Rot{},
}

type CalculatorHandler struct{}

func (CalculatorHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	log.Print("[INFO] " + req.Method + " " + req.URL.Path + " " + req.URL.RawQuery)

	if req.Method != "GET" || req.URL.Path != "/result" {
		e := ErrorMessage{Code: 404, Error: "Not Found"}
		r, _ := json.Marshal(e)
		resp.WriteHeader(404)
		resp.Write(r)
		return
	}

	// Gets the raw query and parse it. The `req.URL.Query()` function is encoding it, making it unusable
	// First, it removes the trailing "op=", and then remove any spaces
	operation := strings.TrimLeft(req.URL.RawQuery, "op=")
	operation = strings.ReplaceAll(operation, " ", "")
	operation = strings.ReplaceAll(operation, "%20", "")

	log.Print("[INFO] Operation query: " + operation)
	log.Print("[INFO] Headers: " + req.Header.Get("Content-Type"))

	if !isOperationValid(operation) {
		c := OperationError{Result: "Invalid expression", Operation: operation}
		r, _ := json.Marshal(c)
		resp.WriteHeader(400)
		resp.Write(r)
		return
	}

	operator := getOperatorFromString(operation)
	mathe := operators[operator]
	params := getOperationParametersFromRawOperation(operation, operator)

	result, calcErr := mathe.Calculate(params)

	if calcErr != nil {
		calcErrMessage := "Could not calculate correctly. Please review the operation"

		if errors.Is(calcErr, ErrInvalidRoot) {
			calcErrMessage = "Invalid root"
		}

		e := ErrorMessage{Code: 400, Error: calcErrMessage}
		r, _ := json.Marshal(e)
		resp.WriteHeader(400)
		resp.Write(r)
		return
	}

	log.Print("[INFO] Operation result: ", result)

	resultString := strconv.FormatFloat(result, 'g', -1, 64)
	log.Print("[INFO] Operation output: ", resultString)

	c := OperationResult{Result: resultString, Operation: operation}
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
