package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"z-blog-go/api/controller"
	"z-blog-go/db/repo"
)

func Register(r *mux.Router) {
	r.HandleFunc("/", controller.IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/sitemap.xml", controller.SitemapHandler).Methods(http.MethodGet)
	r.HandleFunc("/statistics.html", controller.StatisticsHandler).Methods(http.MethodGet)
	r.HandleFunc("/message-board.html", controller.MessageBoardHandler).Methods(http.MethodGet)
	r.HandleFunc("/topic/{topic}.html", controller.TopicHandler).Methods(http.MethodGet)
	r.HandleFunc("/topics.html", controller.TopicsHandler).Methods(http.MethodGet)
	r.HandleFunc("/p/{id:[0-9]+}.html", controller.PostHandler).Methods(http.MethodGet)
	r.HandleFunc("/post/random", controller.PostRandomHandler).Methods(http.MethodPost)
	r.HandleFunc("/search.html", controller.SearchHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("not found", request.RequestURI)
	})
	r.MethodNotAllowedHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("method not allowed")
	})
	r.Use(loggingMiddleware)
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(assets.FS()))))
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/{cmd}", pprof.Index)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer repo.Record(r, time.Now())
		next.ServeHTTP(w, r)
	})
}
