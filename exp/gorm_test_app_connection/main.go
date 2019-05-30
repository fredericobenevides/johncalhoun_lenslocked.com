package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	var psqlInfo string

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()

	us.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@test.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// user.Email = "michael@scott.com"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }
	// userById, err := us.ByID(user.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userById)

	// userByEmail, err := us.ByEmail(user.Email)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(userByEmail)

	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}

}
