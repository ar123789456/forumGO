package models

import (
	"forum/config"
	"log"
)

type Comment struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	User_id int    `json:"user_id"`
	Post_id int    `json:"post_id"`
}

type CommentParams struct {
	Text    string
	User_id int
	Post_id int
}

func (comment *Comment) CREATE(commentParams CommentParams) (*Comment, error) {
	statement, _ := config.DB.Prepare("INSERT INTO comments(text, id_user, id_post) VALUES(?, ?, ?);")
	result, err := statement.Exec(commentParams.Text, commentParams.User_id, commentParams.Post_id)
	if err == nil {
		id, _ := result.LastInsertId()
		comment.ID = int(id)
		comment.Text = commentParams.Text
		comment.Post_id = commentParams.Post_id
		comment.User_id = commentParams.User_id
		return comment, err
	}
	log.Println("Unable to create comment:", err)
	return comment, err
}
