package controller

import (
	"net/http"
	"z-blog-go/db/repo"
)

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	sitemap, _ := repo.GetSitemap(r)
	w.Write([]byte(sitemap))
}
