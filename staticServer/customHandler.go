package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// CustomHandler
type MessageHandler struct {
	message string
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}
func main() {
	// Handler Struct
	m1 := &MessageHandler{message: "hola"}
	http.Handle("/m1", m1)

	// HandlerFunc  Struct
	m2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mensaje 2")
	})
	http.Handle("/m2", m2)

	mensaje3 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}

	http.HandleFunc("/m3", mensaje3)
	http.HandleFunc("/m4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Mensaje 4! %s", time.Now())
	})
	http.Handle("/m5", messageHandler("net/http is awesome"))
	//Handler logic into a Closure

	log.Println("Listening...")
	http.ListenAndServe(":9090", nil)
}
