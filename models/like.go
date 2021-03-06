package models

import (
	"forum/config"
	"log"
)

type Like struct {
	Post_id int `json:"post_id"`
	User_id int `json:"user_id"`
	Liked   int `json:"liked"`
}

func (like *Like) CREATE(Post_id, User_id int) (*Like, error) {
	statement, _ := config.DB.Prepare("INSERT INTO likes_posts(id_post, id_user, liked) VALUES(?, ?, ?);")
	_, err := statement.Exec(Post_id, User_id, LikeTRUE)
	if err == nil {
		like.Post_id = Post_id
		like.User_id = User_id
		return like, err
	}
	log.Println("Unable to create like:", err)
	return like, err
}

func (like *Like) CREATEDISLIKE(Post_id, User_id int) (*Like, error) {
	statement, _ := config.DB.Prepare("INSERT INTO likes_posts(id_post, id_user) VALUES(?, ?);")
	_, err := statement.Exec(Post_id, User_id)
	if err == nil {
		like.Post_id = Post_id
		like.User_id = User_id
		return like, err
	}
	log.Println("Unable to create like:", err)
	return like, err
}

func (like *Like) DELETE(Post_id, User_id int) error {
	statement, _ := config.DB.Prepare("DELETE FROM likes_posts WHERE id_post = ? AND id_user = ?;")
	_, err := statement.Exec(Post_id, User_id)
	return err
}
func (*Like) GETSCORELIKE(Post_id int) (int, error) {
	rows, err := config.DB.Query("SELECT * FROM likes_posts WHERE id_post = ? AND liked = ?", Post_id, LikeTRUE)
	var like int
	if err == nil {
		for rows.Next() {
			like++
		}
		return like, err
	}
	return like, err
}

func (*Like) GETSCOREDISLIKE(Post_id int) (int, error) {
	rows, err := config.DB.Query("SELECT * FROM likes_posts WHERE id_post = ? AND liked = ?", Post_id, Dislike)
	var like int
	if err == nil {
		for rows.Next() {
			like++
		}
		return like, err
	}
	return like, err
}

func (like *Like) GET(Post_id, User_id int) (*Like, error) {
	rows := config.DB.QueryRow("SELECT * FROM likes_posts WHERE id_post = ? AND id_user = ?", Post_id, User_id)
	err := rows.Scan(&like.Post_id, &like.User_id, &like.Liked)
	return like, err
}
