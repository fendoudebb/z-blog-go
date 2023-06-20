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
	web.ExecuteTemplate(w, "post", attr)
}

func PostRandomHandler(w http.ResponseWriter, r *http.Request) {
	randomPosts, _ := repo.GetRandomPosts()
	task.WebsiteStat.RandomPosts = randomPosts
	Ok(w, randomPosts)
}
