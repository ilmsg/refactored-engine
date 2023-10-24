//go:build ignore

package main

import (
	"fmt"

	"github.com/ilmsg/refactored-engine/database"
)

type User struct {
	ID      uint    `json:"id" gorm:"primarykey"`
	Email   string  `json:"email"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserID    uint
}

func main() {
	db := database.GetMySQLConnection()
	// db.Migrator().DropTable(&User{}, &Profile{})
	db.AutoMigrate(&User{}, &Profile{})

	// get
	var users []User
	db.Model(&User{}).Preload("Profile").Find(&users)

	fmt.Printf("users(%d):\n", len(users))
	for _, user := range users {
		fmt.Printf("user: %+v\n", user.Email)
		fmt.Printf("first_name: %+v\n", user.Profile.FirstName)
		fmt.Printf("last_name: %+v\n", user.Profile.LastName)
	}

	// p1 := Profile{FirstName: "Scott", LastName: "Tiger"}
	// p2 := Profile{FirstName: "John", LastName: "Ripper"}
	// db.Create(&p1)
	// db.Create(&p2)

	// u1 := User{Email: "user1@gmail.com", Profile: p1}
	// u2 := User{Email: "user2@gmail.com", Profile: p2}
	// db.Create(&u1)
	// db.Create(&u2)

}
