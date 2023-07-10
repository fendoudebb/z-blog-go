package controller

import (
	"context"
	"net/http"
	"strconv"
	"z-blog-go/configs"
	"z-blog-go/db/repo"
	"z-blog-go/web"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	keywords := r.URL.Query().Get("keywords")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	search := configs.GetApp().Search
	size := search.Size
	maxlength := search.Maxlength
	var posts []repo.Post
	keywordsLength := len(keywords)
	if keywordsLength > 0 {
		if keywordsLength > maxlength {
			keywords = string([]rune(keywords)[:maxlength])
		}
		posts, _ = repo.GetPostsByKeywords(search.Tsconfig, keywords, page, size)
	}
	r = r.WithContext(context.WithValue(r.Context(), "page", page))
	r = r.WithContext(context.WithValue(r.Context(), "size", size))
	r = r.WithContext(context.WithValue(r.Context(), "keywords", keywords))
	r = r.WithContext(context.WithValue(r.Context(), "maxlength", maxlength))
	attr := NewAttr(r, posts)
	web.ExecuteTemplate(w, "search", attr)
}
