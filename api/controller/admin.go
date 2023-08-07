package controller

import (
	"net/http"
	"z-blog-go/web"
)

func AdminIndexHandler(w http.ResponseWriter, r *http.Request) {
	attr := NewAttr(r, nil)
	web.ExecuteTemplate(w, "admin", attr)
}

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	attr := NewAttr(r, nil)
	web.ExecuteTemplate(w, "admin_login", attr)
}
