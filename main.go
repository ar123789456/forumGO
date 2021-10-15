package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

// ErrorCode ...
type ErrorCode struct {
	Code    int
	Message string
}

// Data ...
type Data struct {
	Data []Post
}

// Post ...
type Post struct {
	ID      int
	Name    string
	Content string
	UserID  int
}

// Tag ...
type Tag struct {
	ID     int
	PostID int
	TagID  int
}

// Comment ...
type Comment struct {
	ID      int
	Content string
	PostID  int
	UserID  int
}

var (
	tmpl *template.Template
	data = &Data{}
	db   *sql.DB
)

func init() {
	db, _ = sql.Open("sqlite3", "forum.db")
	CreateTable(db)
	data = Get()
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {

	http.HandleFunc("/", IndexHandler)
	fmt.Println("starting web-server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// CreateTable creates tables in forum db
func CreateTable(db *sql.DB) {
	// create table if not exists
	sqltable, err := os.ReadFile("db/db.txt")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(sqltable))
	if err != nil {
		panic(err)
	}
}

// IndexHandler is main page handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t := r.FormValue("tname")
	t1 := r.FormValue("comment")
	if t != "" && t1 != "" {
		Add(t, t1)
	}
	data = Get()
	tmpl.ExecuteTemplate(w, "index.html", data)
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	var errorPage ErrorCode
	switch status {
	case http.StatusNotFound:
		errorPage.Code = 404
		errorPage.Message = "Page not found"

	case http.StatusInternalServerError:
		errorPage.Code = 500
		errorPage.Message = "Internal server error"

	case http.StatusBadRequest:
		errorPage.Code = 400
		errorPage.Message = "Bad request"

	}
	w.WriteHeader(status)
	tmpl.ExecuteTemplate(w, "error.html", errorPage)
}

// Add ...
func Add(t, t1 string) {
	stmt := "insert into post (name, content, user_id) values (?, ?, ?)"
	//res, err :=
	db.Exec(stmt, t, t1, 1)
}

// Get ...
func Get() *Data {
	rows, _ := db.Query(`
		select name, content, user_id from post
	`)
	var userid int
	var name string
	var content string
	for rows.Next() {
		rows.Scan(&name, &content, &userid)
		data.Data = append(data.Data, Post{
			Name:    name,
			Content: content,
			UserID:  userid,
		})
	}
	return data
}
