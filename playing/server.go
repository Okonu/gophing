package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	ID      string `json:"id"`
	Address Address
}

type Address struct {
	Street  string `json:"street"`
	Town    string `json:"town"`
	ZipKode string `json:"zipKode"`
}

var db *sql.DB

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		"username", "password", "localhost", 3306, "watu")
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	personType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"person": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, ok := params.Args["id"].(string)
					if !ok {
						return nil, fmt.Errorf("invalid ID")
					}

					var person Person
					err := db.QueryRow("SELECT name, age, address_street, address_town, address_zipKode FROM people WHERE id=?", id).Scan(&person.Name, &person.Age, &person.Address.Street, &person.Address.Town, &person.Address.ZipKode)
					if err != nil {
						if err == sql.ErrNoRows {
							return nil, fmt.Errorf("person not found")
						}
						return nil, err
					}
					return person, nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPerson": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"street": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"town": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"zipKode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					name, _ := params.Args["name"].(string)
					age, _ := params.Args["age"].(int)
					street, _ := params.Args["street"].(string)
					town, _ := params.Args["town"].(string)
					zipKode, _ := params.Args["zipKode"].(string)

					result, err := db.Exec("INSERT INTO people (name, age, address_street, address_town, address_zipKode) VALUES (?, ?, ?, ?, ?)", name, age, street, town, zipKode)
					if err != nil {
						return nil, err
					}

					id, err := result.LastInsertId()
					if err != nil {
						return nil, err
					}

					var newPerson Person
					newPerson.ID = fmt.Sprintf("%d", id)
					newPerson.Name = name
					newPerson.Age = age
					newPerson.Address.Street = street
					newPerson.Address.Town = town
					newPerson.Address.ZipKode = zipKode
					return newPerson, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	}).Methods("GET", "POST")

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
