package task

import (
	"log"
	"time"
	"z-blog-go/db/repo"
)

type PostStatistic struct {
	Count      int
	RandomPost []repo.RandomPost
}

var PostStat = &PostStatistic{}

func Start() {
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			postsCount, _ := repo.CountPosts()
			PostStat.Count = postsCount
			posts, _ := repo.GetRandomPosts()
			PostStat.RandomPost = posts
			log.Println("PostStat", PostStat, posts)
			<-ticker.C
		}
	}()
}
