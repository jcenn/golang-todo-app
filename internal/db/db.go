package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Database *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	db_name  = "go_example_todo_db"
)

func ConnectToDB() *sql.DB{
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

