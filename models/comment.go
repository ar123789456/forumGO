package models

import (
	"forum/config"
	"log"
)

type Comment struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	User    string `json:"user"`
	User_id int    `json:"user_id"`
	Post_id int    `json:"post_id"`
	Like    int
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

func (comment *Comment) GET(id int) ([]Comment, error) {
	row, err := config.DB.Query("SELECT * FROM comments WHERE id_post=?", id)
	var comments []Comment

	if err == nil {
		for row.Next() {
			var currentComment Comment
			var user User
			row.Scan(
				&currentComment.ID,
				&currentComment.Text,
				&currentComment.User_id,
				&currentComment.Post_id,
			)

			_, err = user.GetUser(currentComment.User_id)

			if err != nil {
				continue
			}
			currentComment.User = user.Nickname

			comments = append(comments, currentComment)
		}
	}

	return comments, err
}
