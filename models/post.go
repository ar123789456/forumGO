package models

import (
	"forum/config"
	"log"
	"time"
)

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Creat_at  string `json:"creat_at"`
	Update_to string `json:"update_to"`
	User_id   int    `json:"user_id"`
}

type PostParam struct {
	Title       string
	Content     string
	User_id     int
	Category_id int
	// tags        []int
}

func (post *Post) CREATE(userInput PostParam) (*Post, error) {
	statement, _ := config.DB.Prepare("INSERT INTO posts(title, content, create_at, update_at, id_user) VALUES();")
	time := time.Now().String()
	result, err := statement.Exec(userInput.Title, userInput.Content, time, time, userInput.User_id)
	if err == nil {
		id, _ := result.LastInsertId()
		post.Id = int(id)
		post.Content = userInput.Content
		post.Title = userInput.Title
		post.Creat_at = time
		post.Update_to = time
		post.User_id = userInput.User_id
		return post, err
	}
	log.Println("Unable to create post:", err)
	return post, err
}
