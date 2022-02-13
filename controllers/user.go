package controllers

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (*UserController) LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userSession models.UserSession
	if r.Method == http.MethodGet {
		ExecuteLogInTemplate(w, r)
		return
	}
	nick := r.FormValue("Nickname")
	if nick == "" {
		log.Println("form Nickname empty")
		w.WriteHeader(http.StatusBadRequest)
		ExecuteLogInTemplate(w, r)
		return
	}
	_, err := user.FETCH(nick)
	if err != nil {
		log.Println("User not find:", err)
		w.WriteHeader(http.StatusBadRequest)
		ExecuteLogInTemplate(w, r)
		return
	}
	passW := r.FormValue("Password")
	if passW == "" {
		log.Println("form Password empty")
		w.WriteHeader(http.StatusBadRequest)
		ExecuteLogInTemplate(w, r)
		return
	}
	logIn := CheckPasswordHash(passW, user.Password)
	if logIn {
		value := uuid.NewV1().String()
		_, err = userSession.CREATE(value, user.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ExecuteLogInTemplate(w, r)
			return
		}
		cookie := &http.Cookie{
			Name:    "session_token",
			Value:   value,
			Expires: time.Now().Add(120 * time.Minute),
		}
		r.AddCookie(cookie)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("unvalid Password")
	w.WriteHeader(http.StatusBadRequest)
	ExecuteLogInTemplate(w, r)
}

func (*UserController) LogOut(w http.ResponseWriter, r *http.Request) {
	var userSession models.UserSession

	c, err := r.Cookie("session_token")
	if err != nil {
		r.Method = http.MethodGet
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	err = userSession.DELETE(c.Value)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ExecuteLogInTemplate(w http.ResponseWriter, r *http.Request) {
	err := config.Tmpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*UserController) Registration(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userInput models.UserParams
	if r.Method == http.MethodGet {
		err := config.Tmpl.ExecuteTemplate(w, "registration.html", nil)
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
	mail := r.FormValue("Mail")
	if mail == "" {
		log.Println("form Mail empty")

		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	pass := r.FormValue("Password")
	if pass == "" {
		log.Println("form Password empty")

		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	hashPass, err := HashPassword(pass)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	userInput.Email = mail
	userInput.Nickname = nick
	userInput.Password = hashPass

	_, err = user.CREATE(userInput)
	if err != nil {
		log.Println("User not find:", err)
		fmt.Fprint(w, http.StatusBadRequest)
		return
	}
	r.Method = http.MethodGet
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	w.WriteHeader(http.StatusOK)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
