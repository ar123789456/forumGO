package models

import (
	"forum/config"
	"log"
)

type Category struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

func (category *Category) CREATE(Title, Description string) (*Category, error) {
	statement, _ := config.DB.Prepare("INSERT INTO categories(title, description) VALUES(?, ?);")
	result, err := statement.Exec(Title, Description)
	if err == nil {
		id, _ := result.LastInsertId()
		category.Id = int(id)
		category.Title = Title
		category.Discription = Description
		return category, err
	}
	log.Println("Unable to create category:", err)
	return category, err
}
