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
// https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
// https://dev.to/codehakase/building-a-restful-api-with-go
func main() {
    a := App{}
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    a.Run(":8080")
}
