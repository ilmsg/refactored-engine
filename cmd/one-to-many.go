//go:build ignore

package main

import (
	"github.com/google/uuid"
	"github.com/ilmsg/refactored-engine/database"
)

//Paragraph is linked to a story
//A story can have around configurable paragraph
// type Paragraph struct {
//     ID        int
//     StoryID   int
//     Sentences []Sentence `gorm:"ForeignKey:ParagraphID"`
// }

//Sentence are linked to paragraph
//A paragraph can have around configurable paragraphs
// type Sentence struct {
//     ID          uint
//     Value       string
//     Status      bool
//     ParagraphID uint
// }

type Role string

const (
	PO Role = "Product owner"
	M  Role = "Member"
)

type User struct {
	ID    uint      `json:"id" gorm:"primarykey"`
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
	// MemberId uint      `json:"MemberId"`
}

type Member struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UUID      uuid.UUID `json:"uuid"`
	ProjectId uint      `json:"project"`
	User      *User     `json:"user"`
	Role      Role
}

type Task struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UUID      uuid.UUID `json:"uuid"`
	Title     string    `json:"title"`
	ProjectId uint      `json:"project"`
}

type Project struct {
	ID      uint      `json:"id" gorm:"primarykey"`
	UUID    uuid.UUID `json:"uuid"`
	Title   string    `json:"title"`
	Tasks   []*Task   `json:"tasks" gorm:"foreignKey:ProjectId"`
	Members []*Member `json:"members" gorm:"foreignKey:ProjectId"`
}

// UserID  uuid.UUID `json:"_"`

func main() {
	db := database.GetMySQLConnection()
	db.Migrator().DropTable(&User{}, &Member{}, &Task{}, &Project{})
	db.AutoMigrate(&User{}, &Member{}, &Task{}, &Project{})

	u1 := &User{UUID: uuid.New(), Email: "user1@gmail.com"}
	u2 := &User{UUID: uuid.New(), Email: "user2@gmail.com"}
	for u := range []*User{u1, u2} {
		db.Create(u)
	}

	t1 := &Task{UUID: uuid.New(), Title: "task 1"}
	t2 := &Task{UUID: uuid.New(), Title: "task 2"}

	m1 := &Member{UUID: uuid.New(), User: u1, Role: PO}
	m2 := &Member{UUID: uuid.New(), User: u2, Role: M}

	p1 := &Project{UUID: uuid.New(), Title: "Project 1", Tasks: []*Task{t1, t2}, Members: []*Member{m1, m2}}
	db.Create(p1)

	t3 := &Task{UUID: uuid.New(), Title: "task 3"}
	t4 := &Task{UUID: uuid.New(), Title: "task 4"}

	m3 := &Member{UUID: uuid.New(), User: u2, Role: PO}
	m4 := &Member{UUID: uuid.New(), User: u1, Role: M}

	p2 := &Project{UUID: uuid.New(), Title: "Project 2", Tasks: []*Task{t3, t4}, Members: []*Member{m3, m4}}
	db.Create(p2)

	// t1.ProjectId = p1.ID
	// db.Create(t1)

	// db.Create(&Project{ID: uuid.New(), Title: "Project 2", Members: []User{users[1], users[2]}})

	// get
	// var projects []Project
	// db.Model(&Project{}).Preload("Members").Find(&projects)

	// fmt.Printf("projects(%d):\n", len(projects))
	// for _, project := range projects {
	// 	fmt.Printf("title: %+v\n", project.Title)
	// 	fmt.Printf("members(%d):\n", len(project.Members))
	// 	for _, member := range project.Members {
	// 		fmt.Printf("email: %+v\n", member.Email)
	// 	}
	// }
}
