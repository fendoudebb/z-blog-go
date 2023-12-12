package repo

import (
	"fmt"
	"github.com/lib/pq"
	"log"
	"time"
	"z-blog-go/db"
)

type Post struct {
	Id            int
	Uid           int
	Title         string
	Keywords      string
	Description   string
	Topics        []string
	Content       string
	ContentHtml   string
	WordCount     int
	Prop          int
	Status        int
	Pv            int
	LikeCount     int
	CommentCount  int
	CommentStatus int
	CreateTs      time.Time
	UpdateTs      time.Time
}

type RandomPost struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Pv    int    `json:"pv"`
}

func GetPost(id int) (Post, error) {
	var post Post
	row := db.DB.QueryRow("select id, title, keywords, substring(description, 0, 100), topics, content_html, word_count, status, pv, like_count, comment_count, comment_status, create_ts from post where id = $1 limit 1", id)
	if err := row.Scan(&post.Id, &post.Title, &post.Keywords, &post.Description, pq.Array(&post.Topics), &post.ContentHtml, &post.WordCount, &post.Status, &post.Pv, &post.LikeCount, &post.CommentCount, &post.CommentStatus, &post.CreateTs); err != nil {
		return post, fmt.Errorf("GetPost %q: %v", id, err)
	}
	return post, nil
}

func GetPosts(page int, size int) ([]Post, error) {
	var posts []Post
	rows, err := db.DB.Query("select id, title, topics, substring(description, 0, 100), pv, create_ts from post where status = 0 order by id desc limit $1 offset $2", size, (page-1)*size)
	if err != nil {
		return nil, fmt.Errorf("GetPosts: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, pq.Array(&post.Topics), &post.Description, &post.Pv, &post.CreateTs); err != nil {
			return nil, fmt.Errorf("GetPosts: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetRandomPosts() []RandomPost {
	var posts []RandomPost
	rows, _ := db.DB.Query("select id, title, pv from post where status=0 order by random() limit 10")
	defer rows.Close()
	for rows.Next() {
		var post RandomPost
		if err := rows.Scan(&post.Id, &post.Title, &post.Pv); err != nil {
			log.Println("GetRandomPosts err:", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts
}

func GetPostsByTopic(topic string, page int, size int) []Post {
	var posts []Post
	rows, _ := db.DB.Query("select id, title, topics, substring(description, 0, 100), pv, create_ts from post where status=0 and $1=ANY(topics) order by id desc limit $2 offset $3", topic, size, (page-1)*size)
	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, pq.Array(&post.Topics), &post.Description, &post.Pv, &post.CreateTs); err != nil {
			log.Println("GetPostsByTopic err:", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts
}

func GetRankPosts() []Post {
	var posts []Post
	rows, _ := db.DB.Query("select id, title, pv from post where status=0 order by pv desc, id desc limit 5")
	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Pv); err != nil {
			log.Println("GetRankPosts err:", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts
}

func CountPosts() int {
	var count int
	err := db.DB.QueryRow("select count(*) as count from post where status = 0").Scan(&count)
	if err != nil {
		log.Println("CountPosts err:", err)
	}
	return count
}

func CountPostsByTopic(topic string) int {
	var count int
	err := db.DB.QueryRow("select count(*) as count from post where status=0 and $1=ANY(topics)", topic).Scan(&count)
	if err != nil {
		log.Println("CountPostsByTopic err:", err)
	}
	return count
}

func GetPostsByKeywords(tsconfig string, keywords string, page int, size int) ([]Post, error) {
	var posts []Post
	sql := fmt.Sprintf(`
select id, pv, create_ts,
       ts_headline('%s', title, q) as title,
       string_to_array(ts_headline('%s', array_to_string(topics, ','), q), ',') as topics,
       ts_headline('%s', description, q, 'MaxFragments=0, MaxWords=10, MinWords=3, StartSel = <b>, StopSel = </b>') as description
from (select id, title, description, topics, pv, create_ts,
             tsvector_concat(setweight(to_tsvector('%s', title), 'A'), setweight(to_tsvector('%s', description), 'D')) v,
             websearch_to_tsquery('%s', $1) q
      from post where status=0
) temp
where v @@ q order by ts_rank(v, q) desc offset $2 fetch first $3 rows only
`, tsconfig, tsconfig, tsconfig, tsconfig, tsconfig, tsconfig)
	rows, err := db.DB.Query(sql, keywords, (page-1)*size, size)
	if err != nil {
		return nil, fmt.Errorf("GetPostsByKeywords: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Pv, &post.CreateTs, &post.Title, pq.Array(&post.Topics), &post.Description); err != nil {
			return nil, fmt.Errorf("GetPostsByKeywords: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}
