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

func (category *Category) GET(title string) (*Category, error) {
	err := config.DB.QueryRow(
		"SELECT id, title FROM categories WHERE title=?", title).Scan(
		&category.Id, &category.Title)
	return category, err
}

func (category *Category) GETALL() ([]Category, error) {
	rows, err := config.DB.Query("SELECT * FROM categories")
	var categories []Category
	if err == nil {
		for rows.Next() {
			var currentCategory Category
			rows.Scan(
				&currentCategory.Id,
				&currentCategory.Title,
				&currentCategory.Discription,
			)
			categories = append(categories, currentCategory)
		}
		return categories, err
	}
	return categories, err
}
