package main

import (
	"fmt"
	"net/http"
)

// CustomHandler
type MessageHandler struct {
	message string
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)
}

func main() {
	m1 := &MessageHandler{message: "hola"}
	m2 := &MessageHandler{message: "hola 2"}
	http.Handle("/m1", m1)
	http.Handle("/m2", m2)
	http.ListenAndServe(":9090", nil)
}
