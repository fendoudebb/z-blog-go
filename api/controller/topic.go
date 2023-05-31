package controller

import (
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
	"z-blog-go/db/repo"
	"z-blog-go/web"
)

func TopicHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
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
	defer repo.Record(r, time.Now())
	topics, _ := repo.GetTopics()
	attr := NewAttr(r, topics)
	web.ExecuteTemplate(w, "topics", attr)
}
