package task

import (
	"log"
	"time"
	"z-blog-go/db/repo"
)

type websiteStatistic struct {
	PostCount  int
	RandomPost []repo.RandomPost
}

var WebsiteStat = &websiteStatistic{}

func Start() {
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			postsCount, _ := repo.CountPosts()
			WebsiteStat.PostCount = postsCount
			posts, _ := repo.GetRandomPosts()
			WebsiteStat.RandomPost = posts
			log.Println("WebsiteStat", WebsiteStat, posts)
			<-ticker.C
		}
	}()
}
