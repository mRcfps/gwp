package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Post is the post written by some author.
type Post struct {
	ID       int
	Content  string
	Author   string
	Comments []Comment
}

// Comment is comment about a post.
type Comment struct {
	ID      int
	Content string
	Author  string
	Post    *Post
}

// Db is the database used in this program.
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=mRc dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Create creates a new post.
func (c *Comment) Create() (err error) {
	if c.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1, $2, $3) returning id", c.Content, c.Author, c.Post.ID).Scan(&c.ID)
	return
}

// GetPost retrieves a post by its ID field.
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.ID, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// Create stores a post into database.
func (p *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", p.Content, p.Author).Scan(&p.ID)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "mRcfps"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "pftom", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.ID)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
