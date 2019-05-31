package main

import (
	"fmt"

	"lenslocked.com/models"
)

const (
	host     = "192.168.99.100"
	port     = 5432
	user     = "jon"
	password = ""
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()
	us.DestructiveReset()

	user := models.User{
		Name:     "Jon Calhoun",
		Email:    "jon@calhoun.io",
		Password: "jon",
		Remember: "abc123",
	}

	err = us.Create(&user)
	if err != nil {
		panic(err)
	}

	// fmt.Println(user)
	fmt.Printf("%v+", user)

	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	// fmt.Println(user2)
	fmt.Printf("%v+", user2)
}
