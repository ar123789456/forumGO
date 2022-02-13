package controllers

import (
	"fmt"
	"forum/config"
	"forum/models"
	"log"
	"net/http"
	"strconv"
)

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
		allPosts[i].Like, err = like.GETSCORELIKE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		allPosts[i].Dislike, err = like.GETSCOREDISLIKE(post.Id)
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
		allPosts[i].Like, err = like.GETSCORELIKE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		allPosts[i].Dislike, err = like.GETSCOREDISLIKE(post.Id)
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
	var like models.Like
	var likeCom models.LikeComment
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

	if err == nil {
		var err1 error
		post.Like, err1 = like.GETSCORELIKE(post_id)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		post.Dislike, err1 = like.GETSCOREDISLIKE(post_id)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	if err != nil {
		log.Println("Controller/Post GET:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	allComment, err := comment.GET(post_id)

	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for i, comment := range allComment {
		allComment[i].Like, err = likeCom.GETSCORELIKE(comment.ID)
		if err != nil {
			log.Println(err)
		}

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
		allPosts[i].Like, err = like.GETSCORELIKE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		allPosts[i].Dislike, err = like.GETSCOREDISLIKE(post.Id)
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

func (*PostController) GetAllUserPost(w http.ResponseWriter, r *http.Request) {
	var posts models.Post
	var user models.User
	var like models.Like
	var tags models.Tag
	var categories models.Category

	user_idSTR := r.URL.Path[len("/user/"):]

	user_id, err := strconv.Atoi(user_idSTR)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	_, err = user.GetUser(user_id)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	allPosts, err := posts.GETALLUSERPOST(user.Id)

	if err != nil {
		log.Println("Controller/Post:", err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	for i, post := range allPosts {
		allPosts[i].Like, err = like.GETSCORELIKE(post.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		allPosts[i].Dislike, err = like.GETSCOREDISLIKE(post.Id)
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
