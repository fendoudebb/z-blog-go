package controller

import (
	"net/http"
	"time"
	"z-blog-go/db/repo"
	"z-blog-go/web"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
	links, _ := repo.GetLinks()
	attr := NewAttr(r, links)
	web.ExecuteTemplate(w, "links", attr)
}
