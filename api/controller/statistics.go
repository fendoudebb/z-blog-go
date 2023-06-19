package controller

import (
	"net/http"
	"z-blog-go/task"
	"z-blog-go/web"
)

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {
	attr := NewAttr(r, task.WebsiteStat)
	web.ExecuteTemplate(w, "statistics", attr)
}
