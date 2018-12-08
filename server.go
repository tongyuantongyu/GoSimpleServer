package main

import (
    "log"
	"os"
    "net/http"
	"github.com/gorilla/handlers"
)

func main() {
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("www"))))
    err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}