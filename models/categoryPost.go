package models

import (
	"forum/config"
	"log"
)

type CategoryPost struct {
	Id          int
	Category_id int
	Post_id     int
}

func (catPost *CategoryPost) CREATE(Category_id, Post_id int) (*CategoryPost, error) {
	statement, _ := config.DB.Prepare("INSERT INTO category_posts(id_category, id_post) VALUES(?,?);")
	result, err := statement.Exec(Category_id, Post_id)
	if err == nil {
		id, _ := result.LastInsertId()
		catPost.Id = int(id)
		catPost.Category_id = Category_id
		catPost.Post_id = Post_id
		return catPost, err
	}
	log.Println("Unable to create Category_Post:", err)
	return catPost, err
}
