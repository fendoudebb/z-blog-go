package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
	"z-blog-go/db/repo"
	"z-blog-go/task"
	"z-blog-go/web"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	post, _ := repo.GetPost(id)
	attr := NewAttr(r, post)
	attr.Extra = task.WebsiteStat.RandomPost
	web.ExecuteTemplate(w, "post", attr)
}

func PostRandomHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
	posts, _ := repo.GetRandomPosts()
	task.WebsiteStat.RandomPost = posts
	Ok(w, posts)
}
