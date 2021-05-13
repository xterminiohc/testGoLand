package adapter

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func NewUser() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	name := "Edwin"
	email := "codigonota@gmail.com"

	db.Create(&User{Name: name, Email: email})

}

func FindAllUsers() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

}

func FindUserbyName(Name string) User {

	fmt.Println("Params Get IN: " + Name)
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var user User
	db.Where("name = ?", Name).Find(&user)
	fmt.Println("Result Name:" + user.Name)
	fmt.Println("result Email:" + user.Email)

	return user
}
