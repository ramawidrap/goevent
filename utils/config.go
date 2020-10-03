package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "mydatabase"
)

func InitDB() *sql.DB {

	psqlInfo := fmt.Sprintf("postgresql://localhost:5432/mydatabase?sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("failed")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
