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

	_, err = db.Exec(`
	INSERT INTO users(name, email)
	VALUES($1, $2)`, "Jon Calhoun", "jon@calhoun.io")
	if err != nil {
		panic(err)
	}

}
