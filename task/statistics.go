package task

import (
	"github.com/gorilla/mux"
	"github.com/woothee/woothee-go"
	"log"
	"net/http"
	"strconv"
	"time"
	"z-blog-go/db/repo"
	"z-blog-go/helper"
)

type requestInfo struct {
	ip              string
	url             string
	reqMethod       string
	reqParam        string
	userAgent       string
	browser         string
	browserPlatform string
	browserVersion  string
	browserVendor   string
	os              string
	osVersion       string
	referer         string
	searchKeywords  string
	costTime        float64
	invalidReq      bool
	postId          int
}

func Record(r *http.Request, reqTime time.Time) {
	invalidReq, _ := r.Context().Value("invalidReq").(bool) //type cast fail, the value will be false
	info := &requestInfo{
		ip:             helper.GetIp(r),
		url:            r.RequestURI,
		costTime:       time.Since(reqTime).Seconds(), //0.01890375
		reqMethod:      r.Method,
		reqParam:       r.URL.RawQuery,
		userAgent:      r.UserAgent(),
		referer:        r.Referer(),
		searchKeywords: r.URL.Query().Get("keywords"),
		invalidReq:     invalidReq,
	}

	info.postId, _ = strconv.Atoi(mux.Vars(r)["id"])
	go doRecord(info)
}

func doRecord(info *requestInfo) {
	parse, err := woothee.Parse(info.userAgent)
	log.Println("woothee:", parse, err, parse.Category, parse.Os, parse.OsVersion, parse.Name, parse.Version, parse.Vendor, parse.Type)
	ipAddress := repo.QueryAddress("127.0.0.1")
	var reqMethod int
	if info.reqMethod == "GET" {
		reqMethod = 0
	} else if info.reqMethod == "POST" {
		reqMethod = 1
		//TODO req body
	} else if info.reqMethod == "PUT" {
		reqMethod = 2
	} else if info.reqMethod == "DELETE" {
		reqMethod = 3
	} else if info.reqMethod == "OPTION" {
		reqMethod = 4
	} else {
		reqMethod = 5
	}

	log.Println(reqMethod, info, ipAddress)
}
