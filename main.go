package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/varissara-wo/hongjot/api"
	"github.com/varissara-wo/hongjot/migration"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/hongjot?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	// if err := migration.RollbackMigrations(db); err != nil {
	// 	log.Fatal(err)
	// }

	if err := migration.ApplyMigrations(db); err != nil {
		log.Fatal(err)
	}

	e := api.NewServer(db)
	e.Logger.Fatal(e.Start(":8080"))

}
