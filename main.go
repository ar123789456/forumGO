package main

import (
	"forum/config"
	"forum/controllers"
	"forum/middleware"
	"forum/migrations"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	_, err := config.InitializeDB()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = config.ParseTemplates()
	if err != nil {
		log.Println(err)
		return
	}
	migrations.Run()

	var postH controllers.PostController
	var userH controllers.UserController

	mux := http.NewServeMux()
	// router := httprouter.New()

	mux.HandleFunc("/", postH.GetAll)
	mux.Handle("/post/", http.HandlerFunc(postH.GetSinglePost))
	mux.Handle("/post/create", middleware.Authentication(http.HandlerFunc(postH.CreateNewPost)))
	mux.HandleFunc("/tag/", postH.GetAllInTag)
	mux.HandleFunc("/category/", postH.GetAllInCategory)
	mux.HandleFunc("/login", userH.LogIn)
	mux.HandleFunc("/registration", userH.Registration)

	handler := middleware.Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))

	// http.ListenAndServe(":8080", nil)
}
