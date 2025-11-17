package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MilenaBMaciel/MyBookCollection/gen/openapi"
	"github.com/MilenaBMaciel/MyBookCollection/internal/api"
	"github.com/MilenaBMaciel/MyBookCollection/internal/resources/postgres"
	"github.com/jmoiron/sqlx"
	 _ "github.com/lib/pq"
)

func main(){
	fmt.Println("Ola Mundo!")

	dsn := "postgres://username:badpass@localhost:5432/mybookcollection?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	postgresHandler := postgres.New(db)
	handler := api.NewHandler(postgresHandler)

	server, err := openapi.NewServer(handler)
	if err != nil{
		log.Fatal(err)
	}

	if err = http.ListenAndServe(":9090", server); err != nil {
		log.Fatal(err)
	}
}