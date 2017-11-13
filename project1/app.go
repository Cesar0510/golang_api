// app.go

package main

import (
    "database/sql"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

type config struct {
    username string `json:username`
    password string `json:password`
    dbname string `json:dbname`

}

func (a *App) Initialize(user, password, dbname string) { }

func (a *App) Run(addr string) { }

func getPages() []Page {
    raw, err := ioutil.ReadFile("./pages.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Page
    json.Unmarshal(raw, &c)
    return c
}
