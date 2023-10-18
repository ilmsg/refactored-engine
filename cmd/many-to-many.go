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
	for _, user := range users {
		fmt.Printf("email: %+v\n", user.Email)
	}

	// create project
	db.Create(&Project{ID: uuid.New(), Title: "Project 1", Members: []User{users[0]}})
	db.Create(&Project{ID: uuid.New(), Title: "Project 2", Members: []User{users[1], users[2]}})

	// get
	var projects []Project
	db.Model(&Project{}).Preload("Members").Find(&projects)

	fmt.Printf("projects(%d):\n", len(projects))
	for _, project := range projects {
		fmt.Printf("title: %+v\n", project.Title)
		fmt.Printf("members(%d):\n", len(project.Members))
		for _, member := range project.Members {
			fmt.Printf("email: %+v\n", member.Email)
		}
	}
}
