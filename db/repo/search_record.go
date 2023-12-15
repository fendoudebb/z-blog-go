package repo

import (
	"log"
	"z-blog-go/db"
)

type SearchRecord struct {
	Keywords string
	Count    int
}

func GetRankSearchRecords() []SearchRecord {
	var searchRecords []SearchRecord
	rows, _ := db.DB.Query("select keywords, count(keywords) as count from record_search group by keywords order by count desc limit 5")
	defer rows.Close()
	for rows.Next() {
		var searchRecord SearchRecord
		if err := rows.Scan(&searchRecord.Keywords, &searchRecord.Count); err != nil {
			log.Println("GetRankSearchRecords err:", err)
			continue
		}
		searchRecords = append(searchRecords, searchRecord)
	}
	return searchRecords
}
