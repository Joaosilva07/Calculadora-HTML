package main

import (
	"log"
	"net/http"
)

func StartServer() {

	s := &http.Server{
		Addr:    ":8080",
		Handler: CalculatorHandler{},
	}

	log.Fatal(s.ListenAndServe())
}
