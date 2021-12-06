package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", middleware(mainPage))
	http.HandleFunc("/registration", middleware(registrationPage))
	http.HandleFunc("/login", middleware(loginPage))
	http.HandleFunc("/tag", middleware(tagPage))
	http.HandleFunc("/post", middleware(postPage))

	http.ListenAndServe(":8080", nil)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Registration")
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}

func tagPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tag")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post")
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("MiddleWare")
		next(rw, r)
	}
}
