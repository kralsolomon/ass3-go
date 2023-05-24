package main

import (
	"fmt"

	"assignment3/pkg/store/postgres"
)

func main() {
	db, err := postgres.OpenDB("localhost", "5432", "postgres", "password", "dbname")
	if err != nil {
		fmt.Printf("Error connecting to database")
	}
	defer db.Close()
}
