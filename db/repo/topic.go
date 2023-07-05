package repo

import (
	"fmt"
	"z-blog-go/db"
)

type Topic struct {
	Name  string
	Count int
}

func GetTopics() ([]Topic, error) {
	var topics []Topic
	rows, err := db.DB.Query("select count(1) as count, unnest(topics) as name from post group by name order by count desc")
	if err != nil {
		return nil, fmt.Errorf("GetTopics: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var topic Topic
		if err := rows.Scan(&topic.Count, &topic.Name); err != nil {
			return nil, fmt.Errorf("GetTopics: %v", err)
		}
		topics = append(topics, topic)
	}
	return topics, nil
}

func GetRecommendedTopics() ([]string, error) {
	var topics []string
	rows, err := db.DB.Query("select name from topic order by sort")
	if err != nil {
		return nil, fmt.Errorf("GetRecommendedTopics: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var topic string
		if err := rows.Scan(&topic); err != nil {
			return nil, fmt.Errorf("GetRecommendedTopics: %v", err)
		}
		topics = append(topics, topic)
	}
	return topics, nil
}
