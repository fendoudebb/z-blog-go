package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"z-blog-go/db/repo"
	"z-blog-go/task"
	"z-blog-go/web"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	post, _ := repo.GetPost(id)
	attr := NewAttr(r, post)
	attr.Extra = task.PostStat.RandomPost
	web.ExecuteTemplate(w, "post", attr)
}

func PostRandomHandler(w http.ResponseWriter, r *http.Request) {
	posts, _ := repo.GetRandomPosts()
	task.PostStat.RandomPost = posts
	Ok(w, posts)
}
