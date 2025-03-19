package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// main 函数是程序的入口点
func main() {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("无法加载配置: %v", err)
	}

	// 解析命令行参数
	maxResults := flag.Int("max", config.MaxResults, "要处理的最大结果数")
	rssURL := flag.String("rss", "", "单个ArXiv RSS链接（如果提供，将忽略配置中的URL列表）")
	aiProvider := flag.String("ai", config.AIProvider, "AI 提供者 (openai, azure, deepseek)")
	storageProvider := flag.String("storage", config.StorageProvider, "存储提供者 (notion, markdown)")
	runContinuously := flag.Bool("continuous", config.RunContinuously, "持续运行模式")
	// 添加-once标志，与GitHub Actions工作流保持一致
	flag.Bool("once", true, "只运行一次工作流（默认行为，与-continuous相反）")
	generateIndex := flag.Bool("generate-index", false, "只生成Markdown索引，不执行完整工作流")
	flag.Parse()

	// 使用命令行参数更新配置
	config.MaxResults = *maxResults
	config.AIProvider = *aiProvider
	config.StorageProvider = *storageProvider

	// -continuous标志控制是否持续运行
	config.RunContinuously = *runContinuously

	// 如果命令行提供了单一URL，则使用它替换配置中的所有URL
	if *rssURL != "" {
		config.RSSURLs = []string{*rssURL}
	}

	// 如果只需要生成索引
	if *generateIndex && config.StorageProvider == "markdown" {
		log.Println("仅生成Markdown索引...")

		// 创建Markdown存储提供者
		markdownProvider := NewMarkdownStorageProvider(config.OutputDir)

		// 调用生成索引方法
		if err := markdownProvider.GenerateIndex(); err != nil {
			log.Fatalf("生成索引失败: %v", err)
		}

		log.Println("索引生成完成")
		return
	}

	// 运行工作流
	if config.RunContinuously {
		// 持续运行工作流，每天执行一次
		for {
			if err := runWorkflows(config); err != nil {
				log.Printf("工作流执行失败: %v", err)
			}
			// 等待 24 小时
			time.Sleep(24 * time.Hour)
		}
	} else {
		// 仅运行一次工作流
		if err := runWorkflows(config); err != nil {
			log.Fatalf("工作流执行失败: %v", err)
		}
	}
}

// runWorkflows 执行所有数据源的工作流程
func runWorkflows(config *Config) error {
	log.Println("开始执行数据收集工作流...")

	// 步骤1: 创建AI摘要提供者
	log.Printf("使用 %s 作为 AI 提供者", config.AIProvider)
	aiSummarizer, err := NewAISummarizer(config)
	if err != nil {
		return fmt.Errorf("创建 AI 摘要提供者失败: %v", err)
	}

	// 步骤2: 创建存储提供者
	log.Printf("使用 %s 作为存储提供者", config.StorageProvider)
	storageProvider, err := NewStorageProvider(config)
	if err != nil {
		return fmt.Errorf("创建存储提供者失败: %v", err)
	}

	// 步骤3: 创建RSS获取器
	rssFetcher := NewRSSFetcher()

	// 步骤4: 创建去重管理器
	deduplicationManager, err := NewDeduplicationManager(&config.Deduplication)
	if err != nil {
		log.Printf("警告: 创建去重管理器失败: %v，将不进行去重", err)
	}
	// 确保在函数结束时关闭数据库连接
	if deduplicationManager != nil {
		defer deduplicationManager.Close()
	}

	// 步骤5: 处理每个数据源
	totalProcessed := 0
	for i, url := range config.RSSURLs {
		log.Printf("处理数据源 %d/%d: %s", i+1, len(config.RSSURLs), url)

		// 为每个URL单独执行工作流
		processed, err := processRSSURL(url, config.MaxResults, rssFetcher, aiSummarizer, storageProvider, deduplicationManager)
		if err != nil {
			log.Printf("处理数据源失败: %v，继续处理下一个", err)
			continue
		}

		totalProcessed += processed
	}

	log.Printf("工作流执行成功，共处理 %d 个数据源，%d 条数据", len(config.RSSURLs), totalProcessed)
	return nil
}

// processRSSURL 处理单个数据源
func processRSSURL(url string, maxResults int, rssFetcher RSSFetcher, aiSummarizer AISummarizer, storageProvider StorageProvider, deduplicationManager *DeduplicationManager) (int, error) {
	// 步骤1: 获取数据
	log.Printf("正在从 %s 获取数据...", url)
	items, err := rssFetcher.Fetch(url, maxResults)
	if err != nil {
		return 0, fmt.Errorf("获取数据失败: %v", err)
	}
	log.Printf("已获取 %d 条数据", len(items))

	// 步骤2: 处理每条数据
	processedCount := 0
	skippedCount := 0
	for i, item := range items {
		// 去重检查
		if deduplicationManager != nil {
			isDuplicate, similarity, err := deduplicationManager.IsDuplicate(&item)
			if err != nil {
				log.Printf("去重检查时出错: %v, 继续处理该条目", err)
			} else if isDuplicate {
				log.Printf("数据 %s 被识别为重复内容 (相似度: %.2f), 跳过处理", item.Title, similarity)
				skippedCount++
				continue
			}
		}

		log.Printf("处理数据 %d/%d: %s", i+1, len(items), item.Title)

		// 步骤2.1: 使用AI生成摘要
		summary, err := aiSummarizer.Summarize(item.Description)
		if err != nil {
			log.Printf("为数据生成摘要失败: %v, 跳过该条目", err)
			continue
		}
		item.AISummary = summary
		item.SavedAt = time.Now().Format(time.RFC3339)

		// 步骤2.2: 保存数据到存储中
		if err := storageProvider.Save(item); err != nil {
			log.Printf("保存数据失败: %v, 跳过该条目", err)
			continue
		}

		// 添加到去重索引
		if deduplicationManager != nil {
			if err := deduplicationManager.AddItem(&item); err != nil {
				log.Printf("添加到去重索引失败: %v", err)
			}
		}

		log.Printf("数据项 %s 处理完成", item.Title)
		processedCount++
	}

	if skippedCount > 0 {
		log.Printf("由于重复检测，跳过了 %d 条数据", skippedCount)
	}

	return processedCount, nil
}
