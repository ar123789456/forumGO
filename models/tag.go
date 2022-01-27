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

func (tag *Tag) GET(title string) (*Tag, error) {
	err := config.DB.QueryRow(
		"SELECT id, title FROM tags WHERE title=?", title).Scan(
		&tag.Id, &tag.Title)
	return tag, err
}

func (tag *Tag) GETALL() ([]Tag, error) {
	rows, err := config.DB.Query("SELECT * FROM tags")
	var tags []Tag
	if err == nil {
		for rows.Next() {
			var currentTag Tag
			rows.Scan(
				&currentTag.Id,
				&currentTag.Title,
			)
			tags = append(tags, currentTag)
		}
		return tags, err
	}
	return tags, err
}
