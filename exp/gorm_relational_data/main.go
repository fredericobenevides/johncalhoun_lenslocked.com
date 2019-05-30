package main

import (
	"fmt"

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
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
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
	db.AutoMigrate(&Order{})

	var u User
	if err := db.Preload("Orders").First(&u).Error; err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println(u.Orders)

	// var u = User{
	// 	Name:  "blah",
	// 	Email: "blah@blah.com",
	// 	Color: "red",
	// }
	// if err := db.Create(&u).Error; err != nil {
	// 	panic(err)
	// }
	// createOrder(db, u, 1001, "Fake Description #1")
	// createOrder(db, u, 9999, "Fake Description #1")
	// createOrder(db, u, 1001, "Fake Description #2")

}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
