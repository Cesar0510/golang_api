// app.go

package main

import (
    "fmt"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}


func (a *App) Initialize(user, password, dbname string) {
	fmt.Println(user)
  fmt.Println(password)
  fmt.Println(dbname)
}

func (a *App) Run(addr string) {
  fmt.Println(addr)
}
