package controllers

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

type PostController struct{}

func (*PostController) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var params models.PostParam
	var post models.Post
	var tags models.Tag
	var tagPost models.TagPost
	var category models.Category
	var categoryPost models.CategoryPost
	var err error

	if r.Method == http.MethodGet {
		var baseSite struct {
			Tag      []models.Tag
			Category []models.Category
		}
		baseSite.Tag, err = tags.GETALL()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		baseSite.Category, err = category.GETALL()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		config.Tmpl.ExecuteTemplate(w, "addPost.html", baseSite)
		return
	}

	err = params.Parse(r)
	if err != nil {
		log.Println("Controller/ dont pars postParam", err)
		return
	}
	_, err = post.CREATE(params)
	if err != nil {
		fmt.Fprint(w, err)
	}
	_, err = category.GET(params.Category)
	if err != nil {
		post.DELETE(post.Id)
	}
	_, err = categoryPost.CREATE(category.Id, post.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		var baseSite struct {
			Tag      []models.Tag
			Category []models.Category
		}
		baseSite.Tag, err = tags.GETALL()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		baseSite.Category, err = category.GETALL()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		config.Tmpl.ExecuteTemplate(w, "addPost.html", baseSite)
	}
	for _, title := range params.Tags {
		_, err = tags.GET(title)
		if err != nil {
			continue
		}
		_, err = tagPost.CREATE(tags.Id, post.Id)
		if err != nil {
			continue
		}
	}
	w.WriteHeader(http.StatusOK)
	config.Tmpl.ExecuteTemplate(w, "addPost.html", nil)
}

func (*PostController) GetAllInTag(w http.ResponseWriter, r *http.Request) {
}

func (*PostController) GetAllInCategory(w http.ResponseWriter, r *http.Request) {
}

func (*PostController) GetAll(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	allPosts, err := posts.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	err = config.Tmpl.ExecuteTemplate(w, "main.html", allPosts)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) UPDATE(w http.ResponseWriter, r *http.Request) {
	var params models.PostParam
	var post models.Post
	err := params.Parse(r)
	if err != nil {
		log.Println("Controller/Post/Update dont pars postParam", err)
		return
	}
	id, err := strconv.Atoi(r.FormValue("User_id"))
	if err != nil {
		log.Println("Controller/Post/Update dont pars id", err)
		return
	}

	singlePost, err := post.UPDATE(params, id)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	err = config.Tmpl.ExecuteTemplate(w, "main.html", singlePost)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) DELETE(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	id, err := strconv.Atoi(r.FormValue("User_id"))
	if err != nil {
		log.Println("Controller/Post/Delete dont pars id", err)
		return
	}
	err = post.DELETE(id)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	err = config.Tmpl.ExecuteTemplate(w, "main.html", nil)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}
