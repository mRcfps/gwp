package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON data:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
