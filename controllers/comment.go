package controllers

import (
	"fmt"
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

func (*PostController) LikeComment(w http.ResponseWriter, r *http.Request) {
	var like models.LikeComment
	// var user models.User
	var userSession models.UserSession

	if r.URL.Path[len("/comment/like/"):] == "" || r.Method != http.MethodPost {
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
		post_id, err := strconv.Atoi(r.URL.Path[len("/comment/like/"):])
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		like.GET(post_id, user_id)

		if post_id == like.Comment_id && user_id == like.User_id {
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

func (*PostController) DisLikeComment(w http.ResponseWriter, r *http.Request) {
	var like models.LikeComment
	// var user models.User
	var userSession models.UserSession

	if r.URL.Path[len("/comment/dislike/"):] == "" || r.Method != http.MethodPost {
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
		post_id, err := strconv.Atoi(r.URL.Path[len("/comment/dislike/"):])
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		like.GET(post_id, user_id)

		if post_id == like.Comment_id && user_id == like.User_id {
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

func (*PostController) Comment(w http.ResponseWriter, r *http.Request) {
	// var user models.User
	var userSession models.UserSession
	var user_id int
	var commParams models.CommentParams
	var comment models.Comment

	if r.Method == http.MethodPost {
		com := r.FormValue("Comment")
		if com != "" {
			post_id, err := strconv.Atoi(r.URL.Path[len("/post/comment/"):])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)

				log.Println(err)
				return
			}
			c, err := r.Cookie("session_token")
			if err == nil {
				user_id, err = userSession.GetUserId(c.Value)
				if err == nil {
					commParams.Text = com
					commParams.Post_id = post_id
					commParams.User_id = user_id
					_, err := comment.CREATE(commParams)
					if err != nil {
						http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
						log.Println(err)
						return
					}
				}
			} else {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
				w.WriteHeader(http.StatusUnauthorized)
				log.Println(err)
				return
			}

		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}
