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

func (*PostController) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var user models.User
	var comment models.Comment
	var commParams models.CommentParams
	var post_id int
	var user_id int

	if r.URL.Path[6:] != "" {
		var err error
		post_id, err = strconv.Atoi(r.URL.Path[6:])
		if err != nil {
			http.NotFound(w, r)
			return
		}
	}
	if r.Method == http.MethodPost {
		com := r.FormValue("Comment")
		if com != "" {
			c, err := r.Cookie("session_token")
			if err == nil {
				user_id, err = user.GetUserId(c.Value)
				if err == nil {
					commParams.Text = com
					commParams.Post_id = post_id
					commParams.User_id = user_id
					_, err := comment.CREATE(commParams)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	_, err := post.GET(post_id)

	if err != nil {
		log.Println("Controller/Post GET:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	allComment, err := comment.GET(post_id)

	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var allInfo struct {
		AllComment []models.Comment
		Post       models.Post
	}
	allInfo.AllComment = allComment
	allInfo.Post = post

	err = config.Tmpl.ExecuteTemplate(w, "singlePost.html", allInfo)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) GetAll(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	var like models.Like
	allPosts, err := posts.GETALL()

	if r.Method == http.MethodPost {

		post_id, err := strconv.Atoi(r.FormValue("Post_id"))
		if err == nil {
			_, err = like.CREATE(post_id, 1)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

			}
		}
	}

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	for i, post := range allPosts {
		allPosts[i].Like, err = like.GET(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

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
