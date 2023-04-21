package main

import (
	"fmt"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
	ID      string  `json:"id"`
}

type Address struct {
	Street  string `json:"street"`
	Town    string `json:"town"`
	ZipCode string `json:"zipCode"`
}

func main() {
	var p Person
	p.Name = "Jomada"
	p.Age = 100
	p.ID = "1456"
	p.Address = Address{
		Town:    "Kilifi",
		ZipCode: "2014957",
	}

	fmt.Println("%v, %T\n", p, p)
}
