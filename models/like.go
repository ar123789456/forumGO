package models

import (
	"forum/config"
	"log"
)

type Like struct {
	Id      int `json:"id"`
	Post_id int `json:"post_id"`
	User_id int `json:"user_id"`
}

func (like *Like) CREATE(Post_id, User_id int) (*Like, error) {
	statement, _ := config.DB.Prepare("INSERT INTO likes_posts(id_post, id_user) VALUES(?, ?);")
	result, err := statement.Exec(Post_id, User_id)
	if err == nil {
		id, _ := result.LastInsertId()
		like.Id = int(id)
		like.Post_id = Post_id
		like.User_id = User_id
		return like, err
	}
	log.Println("Unable to create like:", err)
	return like, err
}
