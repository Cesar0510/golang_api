package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("public")))
	// Server
	log.Println("Run server .....")
	http.ListenAndServe(":9090", nil)
}
