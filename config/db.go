package config

import "database/sql"

var DB *sql.DB

func InitializeDB() (*sql.DB, error) {
	return sql.Open("sqlite3", "./forum.db")
}
