package controller

import (
	"net/http"
	"time"
	"z-blog-go/db/repo"
	"z-blog-go/web"
)

func MessageBoardHandler(w http.ResponseWriter, r *http.Request) {
	defer repo.Record(r, time.Now())
	attr := NewAttr(r, nil)
	if attr.Website.Comment.Enabled {
		web.ExecuteTemplate(w, "message_board", attr)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
