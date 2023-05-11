package controller

import (
	"log"
	"net/http"
	"z-blog-go/db/repo"
)

func SitemapHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	log.Println(r.Proto)
	log.Println(r.RequestURI)
	log.Println(r.Header.Get("Host"))
	log.Println(r.Header.Get("X-Forwarded-For"))
	log.Println(r.Header.Get("X-Forwarded-Host"))
	log.Println(r.Header.Get("X-Forwarded-Proto"))

	sitemap, _ := repo.GetSitemap()
	w.Write([]byte(sitemap))
}
