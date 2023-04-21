package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	p.ID = "123"
	p.Address.Town = "Kilifi"
	p.Address.ZipCode = "2014957"

	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS people (
		id VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address_street VARCHAR(255) NOT NULL,
		address_town VARCHAR(255) NOT NULL,
		address_zipCode VARCHAR(255) NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO people (id, name, age, address_street, address_town, address_zipCode) VALUES (?, ?, ?, ?, ?, ?)`,
		p.ID, p.Name, p.Age, p.Address.Street, p.Address.Town, p.Address.ZipKode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data saved to database.")
}
