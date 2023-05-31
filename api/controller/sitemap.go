package controller

import (
	"net/http"
	"time"
	"z-blog-go/db/repo"
)

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
	sitemap, _ := repo.GetSitemap(r)
	w.Write([]byte(sitemap))
}
