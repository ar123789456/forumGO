package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var allUser []User

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", middleware(mainPage))
	http.HandleFunc("/registration", middleware(registrationPage))
	http.HandleFunc("/login", middleware(loginPage))
	http.HandleFunc("/tag", middleware(tagPage))
	http.HandleFunc("/post", middleware(postPage))

	http.ListenAndServe(":8080", nil)
}

type User struct {
	id       string
	nickname string
	email    string
	password string
}

func (self *User) initUser(id, nickname, email, password string) {
	self.nickname = nickname
	self.email = email
	self.password = password
	self.id = id
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/registration" {
		log.Println("registration/404")
		return
	}
	switch r.Method {
	case "POST":
		nickname := r.FormValue("Nickname")
		email := r.FormValue("Mail")
		password := r.FormValue("Password")
		hash, _ := HashPassword(password)
		id := uuid.NewV1().String()
		var NewUser User
		NewUser.initUser(id, nickname, email, hash)
		allUser = append(allUser, NewUser)
		fmt.Fprint(w, "Congraits")
	default:
		err := tmpl.ExecuteTemplate(w, "registration.html", nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		log.Println("registration/404")
		return
	}
	switch r.Method {
	case "POST":
		name := r.FormValue("Nickname")
		password := r.FormValue("Password")
		user := findUser(allUser, name, password)

		fmt.Fprint(w, user)
	default:
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func tagPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Tag")
}

func postPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post")
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)
	}
}

func findUser(list []User, name, password string) *User {
	for _, v := range list {
		if v.email == name || v.nickname == name {
			if CheckPasswordHash(password, v.password) {
				return &v
			}
			return nil
		}
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
