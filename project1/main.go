// main.go

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
// https://elithrar.github.io/article/testing-http-handlers-go/
// https://www.thepolyglotdeveloper.com/2016/07/create-a-simple-restful-api-with-golang/
// https://dev.to/codehakase/building-a-restful-api-with-go

type Config struct {
		Database struct{
				Driver string `json:driver`
				Username string `json:username`
				Password string `json:password`
				DBname   string `json:dbname`
			} `json:database`
			Host string `json:hots`
			Port string `json:port`
	}


func LoadConfiguration(file string) Config {
    // https://www.thepolyglotdeveloper.com/2017/04/load-json-configuration-file-golang-application/
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    json.NewDecoder(configFile).Decode(&config)
    return config
}

func main() {
	a := App{}
	config := LoadConfiguration("config.json")
	a.Initialize(config.Database.Username,config.Database.Password, config.Database.DBname)
	a.Run(config.Port)
}
