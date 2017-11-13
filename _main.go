package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	//	"time"
	"log"
)

// https://dev.to/codehakase/building-a-restful-api-with-go
// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

func main() {

	// Basic router with std lib

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	})

	fmt.Println("Server running......")
	http.ListenAndServe(":9090", nil)
}
