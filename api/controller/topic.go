package controller

import (
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"z-blog-go/db/repo"
	"z-blog-go/task"
	"z-blog-go/web"
)

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic, _ := vars["topic"]
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 || size > 20 {
		size = 20
	}
	posts, _ := repo.GetPostsByTopic(topic, page, size)
	log.Println("posts", posts)
	count, _ := repo.CountPostsByTopic(topic)
	attr := NewPaginationAttr(r, posts, page, int(math.Ceil(float64(count)/float64(size))))
	web.ExecuteTemplate(w, "topic", attr)
}

func TopicsHandler(w http.ResponseWriter, r *http.Request) {
	attr := NewAttr(r, task.WebsiteStat.Topics)
	web.ExecuteTemplate(w, "topics", attr)
}
