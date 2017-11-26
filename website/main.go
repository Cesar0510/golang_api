package main

import (
	"fmt"

	"github.com/cesar0510/golang_api/website/models"
)

var lista []int
var data map[string]string

func main() {

	tuple := [4]int{2, 4, 5, 6}
	data := make(map[string]string)
	data["nombre"] = "Cesar"
	cesar := models.User{Name: "Cesar", LastName: "Herdenez"}

	fmt.Println(tuple)
	fmt.Println(cesar)
	fmt.Println(data)

}
