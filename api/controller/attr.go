package controller

import (
	"net/http"
	"time"
	"z-blog-go/configs"
)

type Attr struct {
	Website configs.Website
	Req     *http.Request
	Time    time.Time
	Data    any
	Extra   any
}

type PaginationAttr struct {
	Website     configs.Website
	Req         *http.Request
	Time        time.Time
	Data        any
	CurrentPage int
	SumPage     int
}

func NewAttr(r *http.Request, data any) Attr {
	return Attr{
		Website: configs.GetApp().Website,
		Req:     r,
		Data:    data,
		Time:    time.Now(),
	}
}

func NewPaginationAttr(r *http.Request, data any, currentPage, sumPage int) PaginationAttr {
	return PaginationAttr{
		Website:     configs.GetApp().Website,
		Req:         r,
		Data:        data,
		Time:        time.Now(),
		CurrentPage: currentPage,
		SumPage:     sumPage,
	}
}
