package models

import (
	"forum/config"
	"log"
	"net/http"
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

func (param *PostParam) Parse(r *http.Request) error {
	err := r.ParseForm()
	param.Title = r.FormValue("Title")
	param.Content = r.FormValue("Content")
	return err
}

func (post *Post) CREATE(userInput PostParam) (*Post, error) {
	statement, _ := config.DB.Prepare("INSERT INTO posts(title, content, create_at, update_at, id_user) VALUES(?, ?, ?, ?, ?);")
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

func (post *Post) UPDATE(userInput PostParam, id int) (*Post, error) {
	// ToDo add category_id
	statement, _ := config.DB.Prepare("UPDATE posts SET title = ?, content = ?, update_at = ? WHERE id = ?;")
	time := time.Now().String()
	_, err := statement.Exec(userInput.Title, userInput.Content, time, id)
	if err == nil {
		post.GET(id)
		return post, err
	}
	log.Println("Unable to create post:", err)
	return post, err
}

func (*Post) DELETE(id int) error {
	statement, _ := config.DB.Prepare("DELETE FROM posts WHERE id=?")
	_, err := statement.Exec(id)
	return err
}

func (post *Post) GET(id int) (*Post, error) {
	err := config.DB.QueryRow(
		"SELECT id, title, content, create_at, update_at, id_user FROM posts WHERE id=?", id).Scan(
		&post.Id, &post.Title, &post.Content, &post.Creat_at, &post.Update_to, &post.User_id)
	return post, err
}

func (*Post) GETALL() ([]Post, error) {
	rows, err := config.DB.Query("SELECT * FROM posts")
	var posts []Post
	if err == nil {
		for rows.Next() {
			var currentPost Post
			rows.Scan(
				&currentPost.Id,
				&currentPost.Title,
				&currentPost.Content,
				&currentPost.Creat_at,
				&currentPost.Update_to,
				&currentPost.User_id,
			)
			posts = append(posts, currentPost)
		}
		return posts, err
	}
	return posts, err
}
