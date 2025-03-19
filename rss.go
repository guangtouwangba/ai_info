package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

// DefaultRSSFetcher 是 RSSFetcher 接口的默认实现
type DefaultRSSFetcher struct{}

// NewRSSFetcher 创建一个新的 RSS 获取器
func NewRSSFetcher() RSSFetcher {
	return &DefaultRSSFetcher{}
}

// Fetch 从 URL 获取 RSS 内容并解析成数据项列表
func (f *DefaultRSSFetcher) Fetch(url string, maxResults int) ([]Paper, error) {
	log.Printf("正在从 %s 获取RSS数据", url)

	// 创建 RSS 解析器
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %v", err)
	}

	// 限制结果数量
	resultCount := len(feed.Items)
	if maxResults > 0 && maxResults < resultCount {
		resultCount = maxResults
	}

	// 将 RSS 项转换为数据项
	items := make([]Paper, 0, resultCount)
	for i := 0; i < resultCount; i++ {
		item := feed.Items[i]
		dataItem := Paper{
			Title:       item.Title,
			URL:         item.Link,
			Description: item.Description,
			Summary:     item.Description, // 使用描述作为摘要，将被 AI 处理
			SavedAt:     time.Now().Format(time.RFC3339),
		}
		items = append(items, dataItem)
	}

	log.Printf("已获取 %d 条数据", len(items))
	return items, nil
}

// RSSFetcher 接口定义了获取 RSS 内容的方法
type RSSFetcher interface {
	Fetch(url string, maxResults int) ([]Paper, error)
}

// Paper 表示从数据源获取的信息项
type Paper struct {
	Title       string
	URL         string
	Description string
	Summary     string
	AISummary   string
	SavedAt     string
}
