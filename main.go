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
var allPost []Post
var allTags Tags
var postTag []Post_Tag

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/tag", middleware(tagHandler))
	http.HandleFunc("/post", middleware(postHandler))
	http.HandleFunc("/createpost", middleware(addPostHandler))
	http.HandleFunc("/", mainPage)

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
	fmt.Println("main")
	fmt.Println(len(allPost))
	if r.URL.Path != "/" {
		return
	}
	err := tmpl.ExecuteTemplate(w, "main.html", allPost)
	if err != nil {
		log.Println(err)
	}
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registrationPage")

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
		err := tmpl.ExecuteTemplate(w, "registration.html", nil)
		if err != nil {
			log.Println(err)
		}
	default:
		err := tmpl.ExecuteTemplate(w, "registration.html", nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("loginPage")
	if r.URL.Path != "/login" {
		log.Println("registration/404")
		return
	}
	switch r.Method {
	case "POST":
		name := r.FormValue("Nickname")
		password := r.FormValue("Password")
		user := findUser(allUser, name, password)
		if user == nil {
			r.Method = "GET"
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return
		}
		value := uuid.NewV1().String()
		cookie := &http.Cookie{
			Name:   name,
			Value:  value,
			MaxAge: 300,
		}
		// r.AddCookie(cookie)
		http.SetCookie(w, cookie)
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println(err)
		}
	default:
		err := tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func tagHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tagPage")
	if r.URL.Path[:4] != "/tag" {
		return
	}
	if len(r.URL.Path) < 5 {
		return
	}
	tagPost := tagFindToPost(r.URL.Path[5:])
	err := tmpl.ExecuteTemplate(w, "main.html", tagPost)
	if err != nil {
		log.Println(err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("postPage")
	fmt.Fprint(w, "Post")
}

func addPostHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("addPostHandler")

	if r.URL.Path[:11] != "/createpost" {
		log.Println("createpost/404")
		return
	}
	switch r.Method {
	case "POST":
		title := r.FormValue("Title")
		tag := r.FormValue("Tag")
		content := r.FormValue("Content")
		tAg := allTags.findTag(tag)
		if tAg == nil {
			http.Error(rw, "400 StatusBadRequest", http.StatusBadRequest)
			return
		}
		var post Post
		post.Content = content
		post.Tag = tag
		post.Title = title
		allPost = append(allPost, post)

		var pTag Post_Tag
		pTag.postID = &post
		pTag.TagID = tAg
		postTag = append(postTag, pTag)
		err := tmpl.ExecuteTemplate(rw, "addPost.html", nil)
		if err != nil {
			log.Println(err)
		}
	default:
		err := tmpl.ExecuteTemplate(rw, "addPost.html", nil)
		if err != nil {
			log.Println(err)
		}
	}
}

type Post struct {
	Title   string
	Tag     string
	Content string
}

type Tag struct {
	Name string
}

type Tags []*Tag

func (self *Tags) findTag(tag string) *Tag {
	for _, i := range *self {
		if i.Name == tag {
			return i
		}
	}
	var tAg Tag
	tAg.Name = tag
	allTags = append(allTags, &tAg)
	return &tAg
}

type Post_Tag struct {
	postID *Post
	TagID  *Tag
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if len(r.Cookies()) == 0 {
			r.URL.Path = "/login"
			loginHandler(rw, r)
			return
		}
		next(rw, r)
	}
}

func tagFindToPost(tag string) []Post {
	tagPost := []Post{}
	for _, i := range allPost {
		if i.Tag == tag {
			tagPost = append(tagPost, i)
		}
	}
	return tagPost
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
