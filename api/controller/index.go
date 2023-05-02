package controller

import (
	"math"
	"net/http"
	"strconv"
	"z-blog-go/db/repo"
	"z-blog-go/task"
	"z-blog-go/web"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size < 1 || size > 20 {
		size = 20
	}
	posts, _ := repo.GetPosts(page, size)
	attr := NewPaginationAttr(r, posts, page, int(math.Ceil(float64(task.PostStat.Count)/float64(size))))
	web.ExecuteTemplate(w, "index", attr)
}
