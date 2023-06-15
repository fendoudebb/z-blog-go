package controller

import (
	"net/http"
	"time"
	"z-blog-go/configs"
	"z-blog-go/task"
)

type Attr struct {
	Website     configs.Website
	WebsiteStat *task.WebsiteStatistics
	Req         *http.Request
	Time        time.Time
	Data        any
}

type PaginationAttr struct {
	Website     configs.Website
	WebsiteStat *task.WebsiteStatistics
	Req         *http.Request
	Time        time.Time
	Data        any
	CurrentPage int
	SumPage     int
}

func NewAttr(r *http.Request, data any) Attr {
	return Attr{
		Website:     configs.GetApp().Website,
		WebsiteStat: task.WebsiteStat,
		Req:         r,
		Data:        data,
		Time:        time.Now(),
	}
}

func NewPaginationAttr(r *http.Request, data any, currentPage, sumPage int) PaginationAttr {
	return PaginationAttr{
		Website:     configs.GetApp().Website,
		WebsiteStat: task.WebsiteStat,
		Req:         r,
		Data:        data,
		Time:        time.Now(),
		CurrentPage: currentPage,
		SumPage:     sumPage,
	}
}
