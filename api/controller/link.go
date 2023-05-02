package controller

import (
	"net/http"
	"z-blog-go/db/repo"
	"z-blog-go/web"
)

func LinksHandler(w http.ResponseWriter, r *http.Request) {
	links, _ := repo.GetLinks()
	attr := NewAttr(r, links)
	web.ExecuteTemplate(w, "links", attr)
}
