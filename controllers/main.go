package controllers

import (
	"forum/models"
	"net/http"
)

type PostController struct{}

type UserController struct{}

type allInfoCreate struct {
	User     bool
	Tag      []models.Tag
	Category []models.Category
}

type AllInfo struct {
	User        bool
	UserInfo    models.User
	Allpost     []models.Post
	Alltag      []models.Tag
	Allcategory []models.Category
}

func UAuth(r *http.Request) bool {
	_, err := r.Cookie("session_token")
	return err == nil
}

func UGet(r *http.Request) (models.User, error) {
	var userSession models.UserSession
	var user models.User
	c, err := r.Cookie("session_token")
	if err != nil {
		return user, err
	}
	_, err = userSession.GET(c.Value)
	if err != nil {
		return user, err
	}
	_, err = user.GetUser(userSession.User_id)
	return user, err
}
