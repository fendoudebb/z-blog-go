package repo

import (
	"fmt"
	"z-blog-go/db"
)

func CountPageView() (int, error) {
	var count int
	row := db.DB.QueryRow("select count(*) from record_page_view")
	if err := row.Scan(&count); err != nil {
		return count, fmt.Errorf("CountPageView: %v", err)
	}
	return count, nil
}
