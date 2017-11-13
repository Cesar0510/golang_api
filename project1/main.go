// main.go

package main

import (

    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)
// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// https://elithrar.github.io/article/testing-http-handlers-go/
func main() {
    a := App{}
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    a.Run(":8080")
}
