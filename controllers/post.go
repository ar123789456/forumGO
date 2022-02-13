package controllers

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

func (*PostController) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var params models.PostParam
	var post models.Post
	var tags models.Tag
	var tagPost models.TagPost
	var category models.Category
	var categoryPost models.CategoryPost
	var err error

	var userSession models.UserSession

	if r.Method == http.MethodGet {
		baseSite := GetInfo(w, r)
		config.Tmpl.ExecuteTemplate(w, "addPost.html", baseSite)
		return
	}

	c, err := r.Cookie("session_token")
	if err != nil {
		r.Method = http.MethodGet
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	_, err = userSession.GET(c.Value)
	if err != nil {
		r.Method = http.MethodGet
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	err = params.Parse(r, userSession.User_id)
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
		baseSite := GetInfo(w, r)
		config.Tmpl.ExecuteTemplate(w, "addPost.html", baseSite)
		return
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

func (*PostController) UPDATE(w http.ResponseWriter, r *http.Request) {
	var params models.PostParam
	var post models.Post
	var userSession models.UserSession

	c, err := r.Cookie("session_token")
	if err != nil {
		r.Method = http.MethodGet
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	_, err = userSession.GET(c.Value)
	if err != nil {
		r.Method = http.MethodGet
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	err = params.Parse(r, userSession.User_id)
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

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	// var user models.User
	var userSession models.UserSession

	if r.URL.Path[len("/post/like/"):] == "" || r.Method != http.MethodPost {
		fmt.Fprint(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Header.Get("Referer"))

	if r.Method == http.MethodPost {
		//Check Session cookie
		c, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}
		user_id, err := userSession.GetUserId(c.Value)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}
		post_id, err := strconv.Atoi(r.URL.Path[len("/post/like/"):])
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		like.GET(post_id, user_id)

		if post_id == like.Post_id && user_id == like.User_id {
			err = like.DELETE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			if like.Liked != models.LikeTRUE {
				_, err = like.CREATE(post_id, user_id)
				if err != nil {
					http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
					w.WriteHeader(http.StatusInternalServerError)
					log.Println(err)
					return
				}
			}
		} else {
			_, err = like.CREATE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (*PostController) DisLikePost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	// var user models.User
	var userSession models.UserSession

	if r.URL.Path[len("/post/dislike/"):] == "" || r.Method != http.MethodPost {
		fmt.Fprint(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Println(r.Header.Get("Referer"))

	if r.Method == http.MethodPost {
		//Check Session cookie
		c, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}
		user_id, err := userSession.GetUserId(c.Value)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			return
		}
		post_id, err := strconv.Atoi(r.URL.Path[len("/post/dislike/"):])
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		like.GET(post_id, user_id)

		if post_id == like.Post_id && user_id == like.User_id {
			err = like.DELETE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			if like.Liked != models.Dislike {
				_, err = like.CREATEDISLIKE(post_id, user_id)
				if err != nil {
					http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
					w.WriteHeader(http.StatusInternalServerError)
					log.Println(err)
					return
				}
			}
		} else {
			_, err = like.CREATEDISLIKE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
		}
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func GetInfo(w http.ResponseWriter, r *http.Request) allInfoCreate {
	var category models.Category
	var tags models.Tag
	var baseSite allInfoCreate
	var err error

	baseSite.Tag, err = tags.GETALL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	baseSite.Category, err = category.GETALL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	baseSite.User = UAuth(r)

	return baseSite
}
