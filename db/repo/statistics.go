package repo

import (
	"net/http"
	"time"
)

type requestInfo struct {
	url       string
	costTime  time.Duration
	reqMethod string
	reqParam  string
}

func Record(r *http.Request, reqTime time.Time) {
	info := &requestInfo{
		url:       r.RequestURI,
		costTime:  time.Since(reqTime),
		reqMethod: r.Method,
		reqParam:  r.URL.RawQuery,
	}
	go func() {
		doRecord(info)
	}()
}

func doRecord(info *requestInfo) {
	//ipAddress, _ := QueryAddress("")
}
