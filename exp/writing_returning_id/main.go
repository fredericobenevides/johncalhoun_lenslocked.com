package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.99.100"
	port     = 5432
	user     = "jon"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	var psqlInfo string

	if password == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int

	row := db.QueryRow(`
	INSERT INTO users(name, email)
	VALUES($1, $2)
	RETURNING id`, "Jon Calhoun", "jon@calhoun.io")

	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("id is...", id)
}
