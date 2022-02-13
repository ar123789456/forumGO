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

	mux.HandleFunc("/", postH.GetAll)
	mux.Handle("/post/create", middleware.Authentication(http.HandlerFunc(postH.CreateNewPost)))
	mux.Handle("/post/like/", middleware.Authentication(http.HandlerFunc(postH.LikePost)))
	mux.Handle("/post/dislike/", middleware.Authentication(http.HandlerFunc(postH.DisLikePost)))
	mux.Handle("/post/comment/", middleware.Authentication(http.HandlerFunc(postH.Comment)))
	mux.Handle("/comment/like/", middleware.Authentication(http.HandlerFunc(postH.LikeComment)))
	mux.Handle("/comment/dislike/", middleware.Authentication(http.HandlerFunc(postH.DisLikeComment)))
	mux.Handle("/logout", middleware.Authentication(http.HandlerFunc(userH.LogOut)))

	mux.HandleFunc("/user/", postH.GetAllUserPost)
	mux.HandleFunc("/post/", postH.GetSinglePost)
	mux.HandleFunc("/tag/", postH.GetAllInTag)
	mux.HandleFunc("/category/", postH.GetAllInCategory)
	mux.HandleFunc("/login", userH.LogIn)
	mux.HandleFunc("/registration", userH.Registration)

	handler := middleware.Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))

	// http.ListenAndServe(":8080", nil)
}
