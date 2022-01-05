package models

import (
	"forum/config"
	"log"
)

type Tag struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (tag *Tag) CREATE(title string) (*Tag, error) {
	statement, _ := config.DB.Prepare("INSERT INTO tags(title) VALUES(?);")
	result, err := statement.Exec(title)
	if err == nil {
		id, _ := result.LastInsertId()
		tag.Id = int(id)
		tag.Title = title
		return tag, err
	}
	log.Println("Unable to create tag:", err)
	return tag, err
}
