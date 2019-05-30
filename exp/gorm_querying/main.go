package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "192.168.99.100"
	port     = 5432
	user     = "jon"
	password = ""
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {
	var psqlInfo string

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{})

	var u User
	// newDB := db.Where("email = ?", "blah@blah.com")
	// newDB = newDB.Or("color = ?", "red")
	// newDB = newDB.First(&u)
	db = db.Where("email = ?", "blah@blah.com").First(&u)
	errors := db.GetErrors()
	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}

	// another example
	db = db.Where("email = ?", "blah@blah.com").First(&u)
	if db.RecordNotFound() {
		fmt.Println("No user found!")
	} else if db.Error != nil {
		panic(db.Error)
	}

	// another example
	if err := db.Where("email = ?", "blah@blah.com").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found!")
		case gorm.ErrInvalidSQL:
			fmt.Println("Invalid SQL")
		default:
			panic(err)
		}
	}

	fmt.Println(u)
}
