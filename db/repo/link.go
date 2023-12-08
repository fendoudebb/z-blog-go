package repo

import (
	"log"
	"z-blog-go/db"
)

type Link struct {
	Id             int
	Website        string
	Url            string
	Webmaster      string
	WebmasterEmail string
	Status         int
}

func GetLinks() []Link {
	var links []Link
	rows, _ := db.DB.Query("select website, url from link where status=0 order by sort")
	defer rows.Close()
	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.Website, &link.Url); err != nil {
			log.Println("GetLinks err:", err)
			continue
		}
		links = append(links, link)
	}
	return links
}
