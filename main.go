package main

import (
	"forum/config"
	"forum/controllers"
	"forum/middleware"
	"log"
	"net/http"
)

func main() {
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

	var postH controllers.PostController
	var userH controllers.UserController

	mux := http.NewServeMux()

	mux.HandleFunc("/", postH.GetAll)
	mux.Handle("/post/create", middleware.Authentication(http.HandlerFunc(postH.CreateNewPost)))
	mux.HandleFunc("/login", userH.LogIn)
	mux.HandleFunc("/registration", userH.Registration)

	handler := middleware.Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", handler))

	// http.ListenAndServe(":8080", nil)
}
