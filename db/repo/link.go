package repo

import (
	"fmt"
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

func GetLinks() ([]Link, error) {
	var links []Link
	rows, err := db.DB.Query("select website, url from link where status=0 order by sort")
	if err != nil {
		return nil, fmt.Errorf("GetLinks: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var link Link
		if err := rows.Scan(&link.Website, &link.Url); err != nil {
			return nil, fmt.Errorf("GetLinks: %v", err)
		}
		links = append(links, link)
	}
	return links, nil
}
