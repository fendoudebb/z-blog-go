package repo

import (
	"log"
	"z-blog-go/db"
)

func CountPageView() int {
	var count int
	err := db.DB.QueryRow("select count(*) from record_page_view").Scan(&count)
	if err != nil {
		log.Println("CountPageView:", err)
	}
	return count
}

func InsertPageView() {
	db.DB.Exec(`insert into %s(url, req_method, req_param, ip_id, ua, browser, browser_platform, browser_version, browser_vendor, os, os_version, referer, cost_time, source)
values(%s, %d, %s, %s, %s, '%s', '%s', '%s', '%s', '%s', '%s', %s, %.2f, %d)`)
}
