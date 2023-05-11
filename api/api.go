package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"z-blog-go/api/controller"
	"z-blog-go/assets"
)

func Register(r *mux.Router) {
	r.HandleFunc("/", controller.IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/sitemap.xml", controller.SitemapHandler).Methods(http.MethodGet)
	r.HandleFunc("/message-board.html", controller.MessageBoardHandler).Methods(http.MethodGet)
	r.HandleFunc("/topic/{topic}.html", controller.TopicHandler).Methods(http.MethodGet)
	r.HandleFunc("/topics.html", controller.TopicsHandler).Methods(http.MethodGet)
	r.HandleFunc("/links.html", controller.LinksHandler).Methods(http.MethodGet)
	r.HandleFunc("/p/{id:[0-9]+}.html", controller.PostHandler).Methods(http.MethodGet)
	r.HandleFunc("/post/random", controller.PostRandomHandler).Methods(http.MethodPost)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(assets.FS()))))
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/{cmd}", pprof.Index)
}
