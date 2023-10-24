//go:build ignore

package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ilmsg/refactored-engine/database"
)

type User struct {
	ID    uuid.UUID `json:"id" gorm:"primarykey"`
	Email string    `json:"email"`
}

type Project struct {
	ID    uuid.UUID `json:"id" gorm:"primarykey"`
	Title string    `json:"title"`

	Members []User `json:"members" gorm:"many2many:project_users;"`
}

func main() {
	db := database.GetMySQLConnection()
	db.Migrator().DropTable(&User{}, &Project{})
	db.AutoMigrate(&User{}, &Project{})

	// create users
	db.Create(&User{ID: uuid.New(), Email: "user1@gmail.com"})
	db.Create(&User{ID: uuid.New(), Email: "user2@gmail.com"})
	db.Create(&User{ID: uuid.New(), Email: "user3@gmail.com"})

	var users []User
	db.Find(&users)
	fmt.Printf("users(%d):\n", len(users))
	fmt.Println("-----------------------------------")
	for _, user := range users {
		fmt.Printf("email: %+v\n", user.Email)
	}

	// create project
	db.Create(&Project{ID: uuid.New(), Title: "Project refactored", Members: []User{users[0]}})
	db.Create(&Project{ID: uuid.New(), Title: "Project engine", Members: []User{users[1], users[2]}})

	// get
	var projects []Project
	db.Model(&Project{}).Preload("Members").Find(&projects)

	fmt.Printf("\n\nprojects(%d):\n", len(projects))
	fmt.Println("-----------------------------------")
	for _, project := range projects {
		fmt.Printf("\ntitle: %+v(%d members)\n", project.Title, len(project.Members))
		for key, member := range project.Members {
			fmt.Printf("member%d: %+v\n", key+1, member.Email)
		}
	}
}
