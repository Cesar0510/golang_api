package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

var noteStore = make(map[string]Note)
var id int = 0

func logger(r *http.Request) { log.Println(r) }

//HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	logger(r)
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	// Decode the incoming Note json
	logger(r)
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func main() {
	// New ServerMux
	router := mux.NewRouter()
	router.StrictSlash(false)

	//EndPoints
	router.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	router.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	//router.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	//router.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	// Server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server listening in port :8080")
	server.ListenAndServe()
}
