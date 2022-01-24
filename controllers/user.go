package controllers

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (*UserController) LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Method == http.MethodGet {
		err := config.Tmpl.ExecuteTemplate(w, "login", nil)
		if err != nil {
			fmt.Fprint(w, http.StatusInternalServerError)
		}
		return
	}
	nick := r.FormValue("Nickname")
	if nick == "" {
		log.Println("form Nickname empty")

		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	_, err := user.FETCH(nick)
	if err != nil {
		log.Println("User not find:", err)
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	passW := r.FormValue("Password")
	if passW == "" {
		log.Println("form Password empty")
		fmt.Fprint(w, http.StatusBadRequest)

		return
	}
	logIn := CheckPasswordHash(passW, user.Password)
	if logIn {
		value := uuid.NewV1().String()
		cookie := &http.Cookie{
			Name:   nick,
			Value:  value,
			MaxAge: 300,
		}
		r.AddCookie(cookie)
		http.SetCookie(w, cookie)
		err = config.Tmpl.ExecuteTemplate(w, "loqin.html", nil)
		if err != nil {
			log.Println(err)
			fmt.Fprint(w, err)
		}
		return
	}
	log.Println("unvalid Password")
	fmt.Fprint(w, http.StatusBadRequest)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
