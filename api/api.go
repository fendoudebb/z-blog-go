package api

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"z-blog-go/api/controller"
	"z-blog-go/assets"
	"z-blog-go/task"
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

	r.HandleFunc("/admin/index", controller.AdminIndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/admin/login", controller.AdminLoginHandler).Methods(http.MethodGet)

	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.FS(assets.FS()))))

	//r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	//http.ServeFile(w, r, "./build/index.html")
	//	log.Println("static not found", r.RequestURI)
	//})

	r.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("not found", request.RequestURI)
		request = request.WithContext(context.WithValue(request.Context(), "invalidReq", true))
		//defer task.Record(request, time.Now())
	})
	r.MethodNotAllowedHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("method not allowed")
		request = request.WithContext(context.WithValue(request.Context(), "invalidReq", true))
	})
	r.Use(loggingMiddleware)
	//r.HandleFunc("/debug/pprof/", pprof.Index)
	//r.HandleFunc("/debug/pprof/{cmd}", pprof.Index)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer task.Record(r, time.Now())
		next.ServeHTTP(w, r)
	})
}
