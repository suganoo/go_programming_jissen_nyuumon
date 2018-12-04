package data

import (
	"fmt"
)

var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	PostDeleteAll()
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
	fmt.Println("end setup")
}
