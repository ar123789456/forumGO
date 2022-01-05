package models

import (
	"forum/config"
	"log"
)

type TagPost struct {
	Id      int
	Tag_id  int
	Post_id int
}

func (tag_post *TagPost) CREATE(id_tag, id_post int) (*TagPost, error) {
	statement, _ := config.DB.Prepare("INSERT INTO tag_posts(id_tag, id_post) VALUES(?, ?);")
	result, err := statement.Exec(id_tag, id_post)
	if err == nil {
		id, _ := result.LastInsertId()
		tag_post.Id = int(id)
		tag_post.Tag_id = id_tag
		tag_post.Post_id = id_post
		return tag_post, err
	}
	log.Println("Unable to create tag_post:", err)
	return tag_post, err
}
