package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("HTTP server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
