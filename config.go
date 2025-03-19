package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config 结构体包含应用程序的所有配置
type Config struct {
	// RSS 相关配置
	RSSURLs    []string
	MaxResults int

	// AI 提供者配置
	AIProvider string

	// OpenAI 配置
	OpenAIAPIKey string

	// Azure OpenAI 配置
	AzureOpenAIAPIKey   string
	AzureOpenAIEndpoint string

	// DeepSeek 配置
	DeepSeekAPIKey string

	// 存储提供者配置
	StorageProvider string

	// Notion 存储配置
	NotionAPIKey     string
	NotionDatabaseID string

	// Markdown 存储配置
	OutputDir string

	// 运行模式
	RunContinuously bool

	// 去重配置
	Deduplication DuplicationConfig

	// AWS 配置
	AWSRegion        string // AWS区域
	AWSProfile       string // AWS配置文件
	AWSDynamoDBTable string // DynamoDB表名
}

// DuplicationConfig 定义去重配置
type DuplicationConfig struct {
	Enabled             bool     // 是否启用去重
	Strategies          []string // 启用的策略列表
	SimilarityThreshold float64  // 相似度阈值(0.0-1.0)
	CacheSize           int      // 缓存大小
	DatabasePath        string   // 数据库路径
	DatabaseType        string   // 数据库类型：sqlite, mysql, postgresql
	AWSRegion           string   // AWS区域
	AWSProfile          string   // AWS配置文件
	TablePrefix         string   // DynamoDB表前缀
}

// LoadConfig 加载应用程序配置
func LoadConfig() (*Config, error) {
	// 尝试加载 .env 文件，忽略错误（如果文件不存在）
	_ = godotenv.Load()

	// 初始化 Viper
	v := viper.New()

	// 设置配置文件名称和路径
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	// 设置环境变量前缀并自动绑定环境变量
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 定义默认配置
	v.SetDefault("rss.urls", []string{"https://rsshub.app/arxiv/search_query=abs:ai+agent&start=0&max_results=10&sortBy=lastUpdatedDate"})
	v.SetDefault("rss.max_results", 10)
	v.SetDefault("ai.provider", "openai")
	v.SetDefault("storage.provider", "markdown")
	v.SetDefault("storage.output_dir", "papers")
	v.SetDefault("run.continuous", false)

	// 去重配置默认值
	v.SetDefault("deduplication.enabled", true)
	v.SetDefault("deduplication.strategies", []string{"url", "content_hash"})
	v.SetDefault("deduplication.similarity_threshold", 0.85)
	v.SetDefault("deduplication.cache_size", 1000)
	v.SetDefault("deduplication.database_path", "deduplication.db")
	v.SetDefault("deduplication.database_type", "sqlite")
	v.SetDefault("deduplication.aws_region", "")         // 默认为空，会使用主AWS配置
	v.SetDefault("deduplication.aws_profile", "")        // 默认为空，会使用主AWS配置
	v.SetDefault("deduplication.table_prefix", "dedup_") // 默认表前缀

	// AWS配置默认值
	v.SetDefault("aws.region", "us-east-1")
	v.SetDefault("aws.profile", "default")
	v.SetDefault("aws.dynamodb_table", "arxiv_papers")

	// 尝试读取配置文件
	if err := v.ReadInConfig(); err != nil {
		// 如果配置文件不存在，不报错，而是使用默认值和环境变量
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("读取配置文件失败: %w", err)
		}
	}

	// 兼容现有的环境变量命名
	mapEnvToViper(v, "OPENAI_API_KEY", "ai.openai.api_key")
	mapEnvToViper(v, "AZURE_OPENAI_API_KEY", "ai.azure.api_key")
	mapEnvToViper(v, "AZURE_OPENAI_ENDPOINT", "ai.azure.endpoint")
	mapEnvToViper(v, "DEEPSEEK_API_KEY", "ai.deepseek.api_key")
	mapEnvToViper(v, "NOTION_API_KEY", "storage.notion.api_key")
	mapEnvToViper(v, "NOTION_DATABASE_ID", "storage.notion.database_id")
	mapEnvToViper(v, "OUTPUT_DIR", "storage.output_dir")

	// 处理单个RSS URL到数组的转换
	var rssURLs []string
	if v.IsSet("rss.url") {
		// 如果设置了旧的单一URL，将其添加到列表
		singleURL := v.GetString("rss.url")
		if singleURL != "" {
			rssURLs = append(rssURLs, singleURL)
		}
	}

	// 如果设置了新的URL列表，则使用它们
	if v.IsSet("rss.urls") {
		// 获取原始字符串，用于手动解析
		urlsString := v.GetString("rss.urls")

		// 检查是否来自环境变量，这种情况下需要手动分割
		if urlsString != "" && strings.Contains(urlsString, ",") {
			// 手动分割逗号分隔的 URL
			urlsFromString := strings.Split(urlsString, ",")

			// 清理每个 URL（去除空白）
			for i, url := range urlsFromString {
				urlsFromString[i] = strings.TrimSpace(url)
			}

			// 如果已经有旧格式的URL，合并列表
			if len(rssURLs) > 0 {
				rssURLs = append(rssURLs, urlsFromString...)
			} else {
				rssURLs = urlsFromString
			}
		} else {
			// 尝试使用 Viper 的 GetStringSlice，可能适用于某些情况
			urlsFromConfig := v.GetStringSlice("rss.urls")
			if len(urlsFromConfig) > 0 {
				// 如果已经有旧格式的URL，合并列表
				if len(rssURLs) > 0 {
					rssURLs = append(rssURLs, urlsFromConfig...)
				} else {
					rssURLs = urlsFromConfig
				}
			}
		}
	}

	// 如果没有设置任何URL，使用默认值
	if len(rssURLs) == 0 {
		rssURLs = []string{"https://rsshub.app/arxiv/search_query=abs:ai+agent&start=0&max_results=10&sortBy=lastUpdatedDate"}
	}

	// 创建并返回 Config 结构体
	config := &Config{
		RSSURLs:             rssURLs,
		MaxResults:          v.GetInt("rss.max_results"),
		AIProvider:          v.GetString("ai.provider"),
		OpenAIAPIKey:        v.GetString("ai.openai.api_key"),
		AzureOpenAIAPIKey:   v.GetString("ai.azure.api_key"),
		AzureOpenAIEndpoint: v.GetString("ai.azure.endpoint"),
		DeepSeekAPIKey:      v.GetString("ai.deepseek.api_key"),
		StorageProvider:     v.GetString("storage.provider"),
		NotionAPIKey:        v.GetString("storage.notion.api_key"),
		NotionDatabaseID:    v.GetString("storage.notion.database_id"),
		OutputDir:           v.GetString("storage.output_dir"),
		RunContinuously:     v.GetBool("run.continuous"),
		Deduplication: DuplicationConfig{
			Enabled:             v.GetBool("deduplication.enabled"),
			Strategies:          v.GetStringSlice("deduplication.strategies"),
			SimilarityThreshold: v.GetFloat64("deduplication.similarity_threshold"),
			CacheSize:           v.GetInt("deduplication.cache_size"),
			DatabasePath:        v.GetString("deduplication.database_path"),
			DatabaseType:        v.GetString("deduplication.database_type"),
			AWSRegion:           v.GetString("deduplication.aws_region"),
			AWSProfile:          v.GetString("deduplication.aws_profile"),
			TablePrefix:         v.GetString("deduplication.table_prefix"),
		},
		AWSRegion:        v.GetString("aws.region"),
		AWSProfile:       v.GetString("aws.profile"),
		AWSDynamoDBTable: v.GetString("aws.dynamodb_table"),
	}

	// 如果 OutputDir 是相对路径，确保它存在
	if !filepath.IsAbs(config.OutputDir) {
		if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
			return nil, fmt.Errorf("创建输出目录失败: %w", err)
		}
	}

	// 打印解析后的URL列表，用于调试
	log.Printf("已加载 %d 个数据源: %v", len(config.RSSURLs), config.RSSURLs)

	return config, nil
}

// mapEnvToViper 将环境变量映射到 Viper 配置
func mapEnvToViper(v *viper.Viper, envName, configPath string) {
	if val := os.Getenv(envName); val != "" {
		v.Set(configPath, val)
	}
}
