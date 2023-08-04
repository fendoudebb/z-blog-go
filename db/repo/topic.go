package repo

import (
	"log"
	"z-blog-go/db"
)

type Topic struct {
	Name  string
	Count int
}

func GetTopics() []Topic {
	var topics []Topic
	rows, _ := db.DB.Query("select count(1) as count, unnest(topics) as name from post group by name order by count desc")
	defer rows.Close()
	for rows.Next() {
		var topic Topic
		if err := rows.Scan(&topic.Count, &topic.Name); err != nil {
			log.Println("GetTopics err:", err)
			continue
		}
		topics = append(topics, topic)
	}
	return topics
}

func GetRecommendedTopics() []string {
	var topics []string
	rows, _ := db.DB.Query("select name from topic order by sort")
	defer rows.Close()
	for rows.Next() {
		var topic string
		if err := rows.Scan(&topic); err != nil {
			log.Println("GetRecommendedTopics err:", err)
			continue
		}
		topics = append(topics, topic)
	}
	return topics
}
