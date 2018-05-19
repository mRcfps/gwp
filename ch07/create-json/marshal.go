package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Post is post.
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author is author of post.
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment is comment.
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		ID: 1,
		Content: "Hello World!",
		Author: Author{
			ID: 2,
			Name: "mRcfps"
		}
	}
}