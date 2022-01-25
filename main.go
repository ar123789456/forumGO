package main

import (
	"fmt"
	"forum/config"
	"forum/controllers"
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

	http.HandleFunc("/", postH.GetAll)
	http.HandleFunc("/post/create", postH.CreateNewPost)
	http.HandleFunc("/login", userH.LogIn)
	http.HandleFunc("/registration", userH.Registration)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
