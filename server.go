package main

import (
    "log"
    "net/http"
)

func StartServer() {
    s := &http.Server{
        Addr:    "192.168.68.105:8080",
        Handler: calc_Handler{},
    }
    log.Fatal(s.ListenAndServe())
}