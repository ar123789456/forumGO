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

	http.HandleFunc("/", postH.GetAll)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
