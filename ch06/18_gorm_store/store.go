package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id         int
	Content    string
	Author     string
	Comments   []Comment
	CreatedAt  time.Time
}

type Comment struct {
	Id         int
	Content    string
	Author     string
	PostId     int
	CreatedAt  time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=postgres password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Helloooooo", Author: "suganoooooo"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "gooood comment", Author: "Jooooooe"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "suganoooooo").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments)
}
