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

type allInfoCreate struct {
	Tag      []models.Tag
	Category []models.Category
}

func (*PostController) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var params models.PostParam
	var post models.Post
	var tags models.Tag
	var tagPost models.TagPost
	var category models.Category
	var categoryPost models.CategoryPost
	var err error

	if r.Method == http.MethodGet {
		baseSite := GetInfo(w, r)
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

func (*PostController) GetAllInTag(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	var like models.Like
	var tags models.Tag
	var categories models.Category

	tag := r.URL.Path[len("/tag/"):]

	allPosts, err := posts.GETALLINTAG(tag)

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	for i, post := range allPosts {
		allPosts[i].Like, err = like.GETSCORE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	allTag, err := tags.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	allCategory, err := categories.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	var allInfo struct {
		Allpost     []models.Post
		Alltag      []models.Tag
		Allcategory []models.Category
	}
	allInfo.Allpost = allPosts
	allInfo.Alltag = allTag
	allInfo.Allcategory = allCategory

	err = config.Tmpl.ExecuteTemplate(w, "main.html", allInfo)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) GetAllInCategory(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	var like models.Like
	var tags models.Tag
	var categories models.Category

	category := r.URL.Path[len("/category/"):]

	allPosts, err := posts.GETALLINCATEGORY(category)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	for i, post := range allPosts {
		allPosts[i].Like, err = like.GETSCORE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	allTag, err := tags.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	allCategory, err := categories.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	var allInfo struct {
		Allpost     []models.Post
		Alltag      []models.Tag
		Allcategory []models.Category
	}
	allInfo.Allpost = allPosts
	allInfo.Alltag = allTag
	allInfo.Allcategory = allCategory

	err = config.Tmpl.ExecuteTemplate(w, "main.html", allInfo)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) GetSinglePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	var comment models.Comment
	var post_id int

	if r.URL.Path[6:] != "" {
		var err error
		post_id, err = strconv.Atoi(r.URL.Path[6:])
		if err != nil {
			http.NotFound(w, r)
			return
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
		BaseSite   allInfoCreate
	}
	allInfo.AllComment = allComment
	allInfo.Post = post
	allInfo.BaseSite = GetInfo(w, r)

	err = config.Tmpl.ExecuteTemplate(w, "singlePost.html", allInfo)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
	}
}

func (*PostController) GetAll(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	var like models.Like
	var tags models.Tag
	var categories models.Category
	allPosts, err := posts.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	for i, post := range allPosts {
		allPosts[i].Like, err = like.GETSCORE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	allTag, err := tags.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	allCategory, err := categories.GETALL()

	if err != nil {
		log.Println("Controller/Post:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	var allInfo struct {
		Allpost     []models.Post
		Alltag      []models.Tag
		Allcategory []models.Category
	}
	allInfo.Allpost = allPosts
	allInfo.Alltag = allTag
	allInfo.Allcategory = allCategory

	err = config.Tmpl.ExecuteTemplate(w, "main.html", allInfo)

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

func (*PostController) LikePost(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	var user models.User
	fmt.Println(r.Header.Get("Referer"))

	if r.Method == http.MethodPost {
		//Check Session cookie
		c, err := r.Cookie("session_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			log.Println(err)
			return
		}
		user_id, err := user.GetUserId(c.Value)
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusNotImplemented)
			log.Println(err)
			return
		}
		post_id, err := strconv.Atoi(r.URL.Path[len("/post/like/"):])
		if err != nil {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusBadRequest)
			log.Println(err)
			return
		}
		like.GET(post_id, user_id)

		if post_id == like.Post_id && user_id == like.User_id {
			err = like.DELETE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusNotImplemented)
				log.Println(err)
				return
			}
		} else {
			_, err = like.CREATE(post_id, user_id)
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusNotImplemented)
				log.Println(err)
				return
			}
		}

	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusMethodNotAllowed)
}

func (*PostController) Comment(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var user_id int
	var commParams models.CommentParams
	var comment models.Comment

	if r.Method == http.MethodPost {
		com := r.FormValue("Comment")
		if com != "" {
			post_id, err := strconv.Atoi(r.URL.Path[len("/post/comment/"):])
			if err != nil {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusBadRequest)
				log.Println(err)
				return
			}
			c, err := r.Cookie("session_token")
			if err == nil {
				user_id, err = user.GetUserId(c.Value)
				if err == nil {
					commParams.Text = com
					commParams.Post_id = post_id
					commParams.User_id = user_id
					_, err := comment.CREATE(commParams)
					if err != nil {
						http.Redirect(w, r, r.Header.Get("Referer"), http.StatusInternalServerError)
						log.Println(err)
						return
					}
				}
			} else {
				http.Redirect(w, r, r.Header.Get("Referer"), http.StatusUnauthorized)
				log.Println(err)
				return
			}

		} else {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusMethodNotAllowed)
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

	return baseSite
}
