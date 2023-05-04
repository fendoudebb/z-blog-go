package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"time"
	"z-blog-go/api"
	"z-blog-go/configs"
	"z-blog-go/db"
	"z-blog-go/task"
	"z-blog-go/web"
)

func main() {
	r := mux.NewRouter()
	api.Register(r)
	http.Handle("/", r)

	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	app := configs.GetApp()

	database, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	task.Start()

	pprofSrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.Pprof.Host, app.Pprof.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%d", app.Server.Host, app.Server.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Fatal(pprofSrv.ListenAndServe())
	}()

	go func() {
		log.Println("serve run in pid:", os.Getpid(), "port:", app.Server.Port)
		watcher, _ := web.ConfigTemplateDir("web")
		if watcher != nil {
			defer watcher.Close()
		}
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), app.Server.GracefulTimeout)
	defer cancel()

	log.Println("shutting down")
	_ = pprofSrv.Shutdown(ctx)
	_ = srv.Shutdown(ctx)
	log.Println("shut down")

	os.Exit(0)
}
