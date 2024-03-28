package database

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dbname := os.Getenv("DBNAME")
	user := os.Getenv("USER")
	connection_string := "postgresql://"+user+"@localhost:26257/"+dbname+"?sslmode=disable"
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Fatalf("Database Connection err: %v", err)
		panic(err)
	}
	return db
}
