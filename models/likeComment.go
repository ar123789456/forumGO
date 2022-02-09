package models

import (
	"forum/config"
	"log"
)

type LikeComment struct {
	Comment_id int `json:"comment_id"`
	User_id    int `json:"user_id"`
}

func (like *LikeComment) CREATE(Comment_id, User_id int) (*LikeComment, error) {
	statement, _ := config.DB.Prepare("INSERT INTO likes_comment(id_comment, id_user) VALUES(?, ?);")
	_, err := statement.Exec(Comment_id, User_id)
	if err == nil {
		like.Comment_id = Comment_id
		like.User_id = User_id
		return like, err
	}
	log.Println("Unable to create like:", err)
	return like, err
}

func (like *LikeComment) DELETE(Comment_id, User_id int) error {
	statement, _ := config.DB.Prepare("DELETE FROM likes_comment WHERE id_comment = ? AND id_user = ?")
	_, err := statement.Exec(Comment_id, User_id)
	return err
}

func (*LikeComment) GETSCORE(Comment_id int) (int, error) {
	rows, err := config.DB.Query("SELECT * FROM likes_comment WHERE id_comment = ? ", Comment_id)
	var like int
	if err == nil {
		for rows.Next() {
			like++
		}
		return like, err
	}
	return like, err
}

func (like *LikeComment) GET(Comment_id, User_id int) (*LikeComment, error) {
	rows := config.DB.QueryRow("SELECT * FROM likes_comment WHERE id_comment = ? AND id_user = ?", Comment_id, User_id)
	err := rows.Scan(&like.Comment_id, &like.User_id)
	return like, err
}
