package controller

import (
	"net/http"
	"z-blog-go/web"
)

func MessageBoardHandler(w http.ResponseWriter, r *http.Request) {
	attr := NewAttr(r, nil)
	if attr.Website.Comment.Enabled {
		web.ExecuteTemplate(w, "message_board", attr)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
