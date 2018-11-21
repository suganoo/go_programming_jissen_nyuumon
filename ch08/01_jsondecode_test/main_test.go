package main

import (
	"testing"
)
type Post struct {
        Id       int       `json:"id"`
        Content  string    `json:"content"`
        Author   Author    `json:"author"`
        Comments []Comment `json:"comments"`
}

type Author struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
}

type Comment struct {
        Id      int    `json:"id"`
        Content string `json:"content"`
        Author  string `json:"author"`
}


func TestDecode(t *testing.T) {
	post, err := decode("post.json")

	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now.")
}
