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
			go func() {
				WebsiteStat.PostCount = repo.CountPosts()
			}()
			go func() {
				WebsiteStat.IpCount = repo.CountIps()
			}()
			go func() {
				WebsiteStat.PageView = repo.CountPageView()
			}()
			go func() {
				WebsiteStat.Links = repo.GetLinks()
			}()
			go func() {
				WebsiteStat.RankPosts = repo.GetRankPosts()
			}()
			go func() {
				WebsiteStat.RandomPosts = repo.GetRandomPosts()
			}()
			go func() {
				WebsiteStat.Topics = repo.GetTopics()
			}()
			go func() {
				WebsiteStat.RecommendedTopics = repo.GetRecommendedTopics()
			}()
			go func() {
				WebsiteStat.RankSearchRecords = repo.GetRankSearchRecords()
			}()
			go repo.QueryUnknownIp()
			<-ticker.C
		}
	}()
}
