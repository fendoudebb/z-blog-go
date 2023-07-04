package repo

import (
	"fmt"
	"z-blog-go/db"
)

type SearchRecord struct {
	Keywords string
	Count    int
}

func GetRankSearchRecords() ([]SearchRecord, error) {
	var searchRecords []SearchRecord
	rows, err := db.DB.Query("select keywords, count(keywords) as count from record_search group by keywords order by count desc limit 5")
	if err != nil {
		return nil, fmt.Errorf("GetRankSearchRecords: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var searchRecord SearchRecord
		if err := rows.Scan(&searchRecord.Keywords, &searchRecord.Count); err != nil {
			return nil, fmt.Errorf("GetRankSearchRecords: %v", err)
		}
		searchRecords = append(searchRecords, searchRecord)
	}
	return searchRecords, nil
}
