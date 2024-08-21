package main

import (
	"database/sql"
	"log"
	"os"
)

func dropTables() {
	connStr := os.Getenv("POSTGRES_CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	dropMidwife := `DROP TABLE IF EXISTS midwife`
	dropMother := `DROP TABLE IF EXISTS mother`

	_, err = db.Query(dropMother)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(dropMidwife)
	if err != nil {
		log.Fatal(err)
	}
}
