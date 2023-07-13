package task

import (
	"time"
	"z-blog-go/db/repo"
)

type WebsiteStatistics struct {
	PostCount         int
	IpCount           int
	PageView          int
	Links             []repo.Link
	RankPosts         []repo.Post
	RandomPosts       []repo.RandomPost
	Topics            []repo.Topic
	RecommendedTopics []string
	RankSearchRecords []repo.SearchRecord
}

var WebsiteStat = &WebsiteStatistics{}

func Start() {
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		for {
			postsCount, _ := repo.CountPosts()
			WebsiteStat.PostCount = postsCount
			ips, _ := repo.CountIps()
			WebsiteStat.IpCount = ips
			pageView, _ := repo.CountPageView()
			WebsiteStat.PageView = pageView
			links, _ := repo.GetLinks()
			WebsiteStat.Links = links
			rankPosts, _ := repo.GetRankPosts()
			WebsiteStat.RankPosts = rankPosts
			randomPosts, _ := repo.GetRandomPosts()
			WebsiteStat.RandomPosts = randomPosts
			topics, _ := repo.GetTopics()
			WebsiteStat.Topics = topics
			recommendedTopics, _ := repo.GetRecommendedTopics()
			WebsiteStat.RecommendedTopics = recommendedTopics
			rankSearchRecords, _ := repo.GetRankSearchRecords()
			WebsiteStat.RankSearchRecords = rankSearchRecords
			<-ticker.C
		}
	}()
}
