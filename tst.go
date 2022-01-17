package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var database_url = "host=localhost port=5432 user=usergo password=0000 dbname=myrestdb sslmode=disable"

func NewdbConnect(DatabaseUrl string) (db *sql.DB, err error) {

	db, err = sql.Open("postgres", DatabaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

type user struct {
	id    int
	email string
	pass  string
}

func main() {
	u := &user{id: 1, email: "asdfg", pass: "asdfghfdsf12344t"}
	db, err := NewdbConnect(database_url)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	if err := db.QueryRow(
		"INSERT INTO users (id,email) VALUES ($1,$2)",
		u.id, u.email,
	).Scan(&u.id); err != nil {
		log.Fatal(err)
	}
}
