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
	row := db.DB.QueryRow("select id, title, keywords, description, topics, content_html, word_count, post_status, pv, like_count, comment_count, comment_status, create_ts from post where id = $1 limit 1", id)
	if err := row.Scan(&post.Id, &post.Title, &post.Keywords, &post.Description, pq.Array(&post.Topics), &post.ContentHtml, &post.WordCount, &post.Status, &post.Pv, &post.LikeCount, &post.CommentCount, &post.CommentStatus, &post.CreateTs); err != nil {
		return post, fmt.Errorf("GetPost %q: %v", id, err)
	}
	return post, nil
}

func GetPosts(page int, size int) ([]Post, error) {
	var posts []Post
	rows, err := db.DB.Query("select id, title, topics, description, pv, create_ts from post where post_status = 0 order by id desc limit $1 offset $2", size, (page-1)*size)
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

func GetRandomPosts() ([]RandomPost, error) {
	var posts []RandomPost
	rows, err := db.DB.Query("select id, title, pv from post where post_status=0 order by random() limit 10")
	if err != nil {
		return nil, fmt.Errorf("GetRandomPosts: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post RandomPost
		if err := rows.Scan(&post.Id, &post.Title, &post.Pv); err != nil {
			return nil, fmt.Errorf("GetRandomPosts: %v", err)
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostsByTopic(topic string, page int, size int) ([]Post, error) {
	var posts []Post
	rows, err := db.DB.Query("select id, title, topics, description, pv, create_ts from post where post_status=0 and $1=ANY(topics) order by id desc limit $2 offset $3", topic, size, (page-1)*size)
	if err != nil {
		return nil, fmt.Errorf("GetPostsByTopic: %v", err)
	}
	log.Println("xxxxxddadas", rows, page, size, topic)
	defer rows.Close()
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, pq.Array(&post.Topics), &post.Description, &post.Pv, &post.CreateTs); err != nil {
			return nil, fmt.Errorf("GetPostsByTopic: %v", err)
		}
		log.Println("xxxxx", post)
		posts = append(posts, post)
	}
	return posts, nil
}

func CountPosts() (int, error) {
	var count int
	row := db.DB.QueryRow("select count(*) as count from post where post_status = 0")
	if err := row.Scan(&count); err != nil {
		return count, fmt.Errorf("CountPosts: %v", err)
	}
	return count, nil
}

func CountPostsByTopic(topic string) (int, error) {
	var count int
	row := db.DB.QueryRow("select count(*) as count from post where post_status=0 and $1=ANY(topics)", topic)
	if err := row.Scan(&count); err != nil {
		return count, fmt.Errorf("CountPostsByTopic: %v", err)
	}
	return count, nil
}
