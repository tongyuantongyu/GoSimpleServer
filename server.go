package main

import (
	"fmt"
    "log"
	"os"
    "net/http"
	"github.com/gorilla/handlers"
)

func main() {
	var port string
	fmt.Print("Type a port (2333 default): ")
	fmt.Scanln(&port)
	if port == "" {port = "2333"}
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("file"))))
    err := http.ListenAndServe(":" + port, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}