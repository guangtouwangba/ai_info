package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	_ "github.com/mattn/go-sqlite3"
	// 如果您需要使用MySQL，请取消下面这行的注释
	// _ "github.com/go-sql-driver/mysql"
)

// DBProvider 定义数据库提供者接口
type DBProvider interface {
	// 初始化数据库，创建必要的表和索引
	Initialize() error

	// 关闭数据库连接
	Close() error

	// URL索引相关方法
	URLExists(url string) (bool, error)
	AddURL(url, itemID string) error

	// 内容哈希索引相关方法
	ContentHashExists(hash string) (bool, error)
	AddContentHash(hash, itemID string) error

	// 标题索引相关方法
	TitleExists(normalizedTitle string) (bool, error)
	AddTitle(normalizedTitle, itemID string) error

	// 内容相似度索引相关方法
	GetLatestContentItems(limit int) ([]ContentItem, error)
	AddContentItem(itemID, content, contentTokens string) error
}

// ContentItem 表示用于相似度比较的内容项
type ContentItem struct {
	ItemID        string
	Content       string
	ContentTokens []string
}

// NewDBProvider 根据配置创建适当的数据库提供者
func NewDBProvider(config *DuplicationConfig) (DBProvider, error) {
	switch config.DatabaseType {
	case "sqlite", "":
		return NewSQLiteProvider(config.DatabasePath)
	case "mysql":
		// 将来实现MySQL提供者
		if strings.HasPrefix(config.DatabasePath, "mysql://") {
			return NewMySQLProvider(config.DatabasePath)
		}
		return nil, fmt.Errorf("MySQL 连接字符串格式错误，应以 mysql:// 开头")
	case "dynamodb":
		// 使用DynamoDB提供者
		return NewDynamoDBProvider(config.TablePrefix, config)
	case "postgresql":
		// 将来实现PostgreSQL提供者
		return nil, fmt.Errorf("PostgreSQL 数据库暂未实现")
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", config.DatabaseType)
	}
}

// SQLiteProvider 实现基于SQLite的数据库提供者
type SQLiteProvider struct {
	db   *sql.DB
	path string
	mu   sync.Mutex
}

// NewSQLiteProvider 创建一个新的SQLite提供者
func NewSQLiteProvider(dbPath string) (*SQLiteProvider, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("无法打开数据库: %v", err)
	}

	provider := &SQLiteProvider{
		db:   db,
		path: dbPath,
	}

	if err := provider.Initialize(); err != nil {
		db.Close()
		return nil, err
	}

	return provider, nil
}

// Initialize 实现DBProvider接口，初始化数据库
func (p *SQLiteProvider) Initialize() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 创建URL索引表
	_, err := p.db.Exec(`CREATE TABLE IF NOT EXISTS url_index (
		url TEXT PRIMARY KEY,
		item_id TEXT,
		created_at TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("创建URL索引表失败: %v", err)
	}

	// 创建内容哈希索引表
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS content_hash_index (
		content_hash TEXT PRIMARY KEY,
		item_id TEXT,
		created_at TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("创建内容哈希索引表失败: %v", err)
	}

	// 创建标题索引表
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS title_index (
		normalized_title TEXT PRIMARY KEY,
		item_id TEXT,
		created_at TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("创建标题索引表失败: %v", err)
	}

	// 创建内容索引表(用于相似度检查)
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS content_index (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id TEXT,
		content TEXT,
		content_tokens TEXT,
		created_at TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("创建内容索引表失败: %v", err)
	}

	// 创建相关索引
	_, err = p.db.Exec(`CREATE INDEX IF NOT EXISTS idx_content_index_item_id ON content_index(item_id)`)
	if err != nil {
		return fmt.Errorf("创建索引失败: %v", err)
	}

	return nil
}

// Close 关闭数据库连接
func (p *SQLiteProvider) Close() error {
	return p.db.Close()
}

// URLExists 检查URL是否存在
func (p *SQLiteProvider) URLExists(url string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM url_index WHERE url = ?", url).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("URL查询失败: %v", err)
	}

	return count > 0, nil
}

// AddURL 添加URL到索引
func (p *SQLiteProvider) AddURL(url, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT OR IGNORE INTO url_index (url, item_id, created_at) VALUES (?, ?, ?)",
		url, itemID, time.Now())

	return err
}

// ContentHashExists 检查内容哈希是否存在
func (p *SQLiteProvider) ContentHashExists(hash string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM content_hash_index WHERE content_hash = ?", hash).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("内容哈希查询失败: %v", err)
	}

	return count > 0, nil
}

// AddContentHash 添加内容哈希到索引
func (p *SQLiteProvider) AddContentHash(hash, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT OR IGNORE INTO content_hash_index (content_hash, item_id, created_at) VALUES (?, ?, ?)",
		hash, itemID, time.Now())

	return err
}

// TitleExists 检查标题是否存在
func (p *SQLiteProvider) TitleExists(normalizedTitle string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM title_index WHERE normalized_title = ?", normalizedTitle).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("标题查询失败: %v", err)
	}

	return count > 0, nil
}

// AddTitle 添加标题到索引
func (p *SQLiteProvider) AddTitle(normalizedTitle, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT OR IGNORE INTO title_index (normalized_title, item_id, created_at) VALUES (?, ?, ?)",
		normalizedTitle, itemID, time.Now())

	return err
}

// GetLatestContentItems 获取最新的内容项目
func (p *SQLiteProvider) GetLatestContentItems(limit int) ([]ContentItem, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	rows, err := p.db.Query(
		"SELECT item_id, content, content_tokens FROM content_index ORDER BY created_at DESC LIMIT ?",
		limit)
	if err != nil {
		return nil, fmt.Errorf("获取内容项目失败: %v", err)
	}
	defer rows.Close()

	var items []ContentItem
	for rows.Next() {
		var item ContentItem
		var tokensStr string
		if err := rows.Scan(&item.ItemID, &item.Content, &tokensStr); err != nil {
			return nil, fmt.Errorf("解析内容项目失败: %v", err)
		}

		// 解析tokens字符串
		if tokensStr != "" {
			item.ContentTokens = strings.Split(tokensStr, ",")
		}

		items = append(items, item)
	}

	return items, nil
}

// AddContentItem 添加内容项目到索引
func (p *SQLiteProvider) AddContentItem(itemID, content, contentTokens string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT INTO content_index (item_id, content, content_tokens, created_at) VALUES (?, ?, ?, ?)",
		itemID, content, contentTokens, time.Now())

	return err
}

// MySQLProvider 实现基于MySQL的数据库提供者
type MySQLProvider struct {
	db  *sql.DB
	dsn string
	mu  sync.Mutex
}

// NewMySQLProvider 创建一个新的MySQL提供者
func NewMySQLProvider(connString string) (*MySQLProvider, error) {
	// 将 mysql:// 格式的连接字符串转换为 DSN 格式
	// 例如: mysql://user:password@host:port/dbname 转换为 user:password@tcp(host:port)/dbname
	dsn := strings.TrimPrefix(connString, "mysql://")

	// 检查是否需要导入MySQL驱动
	// _ "github.com/go-sql-driver/mysql"
	log.Println("警告: 请确保您的项目已导入MySQL驱动 (github.com/go-sql-driver/mysql)")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("无法连接MySQL数据库: %v", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("MySQL数据库连接测试失败: %v", err)
	}

	provider := &MySQLProvider{
		db:  db,
		dsn: dsn,
	}

	if err := provider.Initialize(); err != nil {
		db.Close()
		return nil, err
	}

	return provider, nil
}

// Initialize 实现DBProvider接口，初始化数据库
func (p *MySQLProvider) Initialize() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 创建URL索引表
	_, err := p.db.Exec(`CREATE TABLE IF NOT EXISTS url_index (
		url VARCHAR(512) PRIMARY KEY,
		item_id VARCHAR(512) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	if err != nil {
		return fmt.Errorf("创建URL索引表失败: %v", err)
	}

	// 创建内容哈希索引表
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS content_hash_index (
		content_hash VARCHAR(64) PRIMARY KEY,
		item_id VARCHAR(512) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	if err != nil {
		return fmt.Errorf("创建内容哈希索引表失败: %v", err)
	}

	// 创建标题索引表
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS title_index (
		normalized_title VARCHAR(512) PRIMARY KEY,
		item_id VARCHAR(512) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	if err != nil {
		return fmt.Errorf("创建标题索引表失败: %v", err)
	}

	// 创建内容索引表(用于相似度检查)
	_, err = p.db.Exec(`CREATE TABLE IF NOT EXISTS content_index (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		item_id VARCHAR(512) NOT NULL,
		content TEXT,
		content_tokens TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		INDEX idx_content_index_item_id (item_id),
		INDEX idx_content_index_created_at (created_at)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	if err != nil {
		return fmt.Errorf("创建内容索引表失败: %v", err)
	}

	return nil
}

// Close 关闭数据库连接
func (p *MySQLProvider) Close() error {
	return p.db.Close()
}

// URLExists 检查URL是否存在
func (p *MySQLProvider) URLExists(url string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM url_index WHERE url = ?", url).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("MySQL URL查询失败: %v", err)
	}

	return count > 0, nil
}

// AddURL 添加URL到索引
func (p *MySQLProvider) AddURL(url, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT IGNORE INTO url_index (url, item_id) VALUES (?, ?)",
		url, itemID)

	return err
}

// ContentHashExists 检查内容哈希是否存在
func (p *MySQLProvider) ContentHashExists(hash string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM content_hash_index WHERE content_hash = ?", hash).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("MySQL内容哈希查询失败: %v", err)
	}

	return count > 0, nil
}

// AddContentHash 添加内容哈希到索引
func (p *MySQLProvider) AddContentHash(hash, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT IGNORE INTO content_hash_index (content_hash, item_id) VALUES (?, ?)",
		hash, itemID)

	return err
}

// TitleExists 检查标题是否存在
func (p *MySQLProvider) TitleExists(normalizedTitle string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var count int
	err := p.db.QueryRow("SELECT COUNT(*) FROM title_index WHERE normalized_title = ?", normalizedTitle).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("MySQL标题查询失败: %v", err)
	}

	return count > 0, nil
}

// AddTitle 添加标题到索引
func (p *MySQLProvider) AddTitle(normalizedTitle, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT IGNORE INTO title_index (normalized_title, item_id) VALUES (?, ?)",
		normalizedTitle, itemID)

	return err
}

// GetLatestContentItems 获取最新的内容项目
func (p *MySQLProvider) GetLatestContentItems(limit int) ([]ContentItem, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	rows, err := p.db.Query(
		"SELECT item_id, content, content_tokens FROM content_index ORDER BY created_at DESC LIMIT ?",
		limit)
	if err != nil {
		return nil, fmt.Errorf("MySQL获取内容项目失败: %v", err)
	}
	defer rows.Close()

	var items []ContentItem
	for rows.Next() {
		var item ContentItem
		var tokensStr string
		if err := rows.Scan(&item.ItemID, &item.Content, &tokensStr); err != nil {
			return nil, fmt.Errorf("MySQL解析内容项目失败: %v", err)
		}

		// 解析tokens字符串
		if tokensStr != "" {
			item.ContentTokens = strings.Split(tokensStr, ",")
		}

		items = append(items, item)
	}

	return items, nil
}

// AddContentItem 添加内容项目到索引
func (p *MySQLProvider) AddContentItem(itemID, content, contentTokens string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.db.Exec(
		"INSERT INTO content_index (item_id, content, content_tokens) VALUES (?, ?, ?)",
		itemID, content, contentTokens)

	return err
}

// Deduplicator 定义去重功能的接口
type Deduplicator interface {
	// IsDuplicate 检查一个数据项是否是重复的
	// 返回 (是否重复, 与现有条目的相似度, 错误)
	IsDuplicate(item *Paper) (bool, float64, error)

	// AddItem 添加一个项目到去重器的索引/缓存中
	AddItem(item *Paper) error

	// Name 返回去重策略的名称，用于日志和调试
	Name() string
}

// DeduplicationManager 管理多个去重策略
type DeduplicationManager struct {
	strategies []Deduplicator
	config     *DuplicationConfig
	dbProvider DBProvider
}

// NewDeduplicationManager 创建一个新的去重管理器
func NewDeduplicationManager(config *DuplicationConfig) (*DeduplicationManager, error) {
	if !config.Enabled {
		log.Println("去重功能已禁用")
		return &DeduplicationManager{
			strategies: []Deduplicator{},
			config:     config,
		}, nil
	}

	log.Printf("初始化去重管理器，启用策略: %v，数据库类型: %s", config.Strategies, config.DatabaseType)

	// 创建数据库提供者
	dbProvider, err := NewDBProvider(config)
	if err != nil {
		return nil, fmt.Errorf("创建数据库提供者失败: %v", err)
	}

	strategies := make([]Deduplicator, 0, len(config.Strategies))

	// 根据配置初始化不同的去重策略
	for _, strategyName := range config.Strategies {
		var strategy Deduplicator
		var err error

		switch strategyName {
		case "url":
			strategy, err = NewURLDeduplicator(config.CacheSize, dbProvider)
		case "content_hash":
			strategy, err = NewContentHashDeduplicator(config.CacheSize, dbProvider)
		case "title":
			strategy, err = NewTitleDeduplicator(config.CacheSize, dbProvider)
		case "similarity":
			strategy, err = NewSimilarityDeduplicator(config.SimilarityThreshold, config.CacheSize, dbProvider)
		default:
			log.Printf("未知的去重策略: %s, 已跳过", strategyName)
			continue
		}

		if err != nil {
			log.Printf("初始化去重策略 %s 失败: %v", strategyName, err)
			continue
		}

		strategies = append(strategies, strategy)
		log.Printf("已启用去重策略: %s", strategy.Name())
	}

	return &DeduplicationManager{
		strategies: strategies,
		config:     config,
		dbProvider: dbProvider,
	}, nil
}

// IsDuplicate 使用所有配置的策略检查是否重复
func (dm *DeduplicationManager) IsDuplicate(item *Paper) (bool, float64, error) {
	if !dm.config.Enabled || len(dm.strategies) == 0 {
		return false, 0, nil
	}

	// 依次使用每个策略检查
	for _, strategy := range dm.strategies {
		isDuplicate, similarity, err := strategy.IsDuplicate(item)
		if err != nil {
			log.Printf("使用策略 %s 检查重复时出错: %v", strategy.Name(), err)
			continue
		}

		if isDuplicate {
			log.Printf("策略 %s 检测到重复，相似度: %.2f", strategy.Name(), similarity)
			return true, similarity, nil
		}
	}

	return false, 0, nil
}

// AddItem 向所有策略添加项目
func (dm *DeduplicationManager) AddItem(item *Paper) error {
	if !dm.config.Enabled || len(dm.strategies) == 0 {
		return nil
	}

	for _, strategy := range dm.strategies {
		if err := strategy.AddItem(item); err != nil {
			log.Printf("添加项目到策略 %s 失败: %v", strategy.Name(), err)
		}
	}

	return nil
}

// Close 关闭数据库连接
func (dm *DeduplicationManager) Close() error {
	if dm.dbProvider != nil {
		return dm.dbProvider.Close()
	}
	return nil
}

// ---------------- URL去重策略 ----------------

// URLDeduplicator 基于URL的去重策略
type URLDeduplicator struct {
	cache      *lru.Cache[string, bool]
	dbProvider DBProvider
}

// NewURLDeduplicator 创建URL去重器
func NewURLDeduplicator(cacheSize int, dbProvider DBProvider) (*URLDeduplicator, error) {
	cache, err := lru.New[string, bool](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("创建缓存失败: %v", err)
	}

	return &URLDeduplicator{
		cache:      cache,
		dbProvider: dbProvider,
	}, nil
}

// IsDuplicate 检查URL是否已存在
func (d *URLDeduplicator) IsDuplicate(item *Paper) (bool, float64, error) {
	// 首先检查缓存
	if _, found := d.cache.Get(item.URL); found {
		return true, 1.0, nil
	}

	// 查询数据库
	exists, err := d.dbProvider.URLExists(item.URL)
	if err != nil {
		return false, 0, err
	}

	if exists {
		// 添加到缓存
		d.cache.Add(item.URL, true)
	}

	return exists, 1.0, nil
}

// AddItem 将URL添加到索引
func (d *URLDeduplicator) AddItem(item *Paper) error {
	// 添加到缓存
	d.cache.Add(item.URL, true)

	// 添加到数据库
	return d.dbProvider.AddURL(item.URL, item.URL)
}

// Name 返回策略名称
func (d *URLDeduplicator) Name() string {
	return "URL去重"
}

// ---------------- 内容哈希去重策略 ----------------

// ContentHashDeduplicator 基于内容哈希的去重策略
type ContentHashDeduplicator struct {
	cache      *lru.Cache[string, bool]
	dbProvider DBProvider
}

// NewContentHashDeduplicator 创建内容哈希去重器
func NewContentHashDeduplicator(cacheSize int, dbProvider DBProvider) (*ContentHashDeduplicator, error) {
	cache, err := lru.New[string, bool](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("创建缓存失败: %v", err)
	}

	return &ContentHashDeduplicator{
		cache:      cache,
		dbProvider: dbProvider,
	}, nil
}

// IsDuplicate 检查内容哈希是否已存在
func (d *ContentHashDeduplicator) IsDuplicate(item *Paper) (bool, float64, error) {
	contentHash := generateContentHash(item.Description)

	// 首先检查缓存
	if _, found := d.cache.Get(contentHash); found {
		return true, 1.0, nil
	}

	// 查询数据库
	exists, err := d.dbProvider.ContentHashExists(contentHash)
	if err != nil {
		return false, 0, err
	}

	if exists {
		// 添加到缓存
		d.cache.Add(contentHash, true)
	}

	return exists, 1.0, nil
}

// AddItem 将内容哈希添加到索引
func (d *ContentHashDeduplicator) AddItem(item *Paper) error {
	contentHash := generateContentHash(item.Description)

	// 添加到缓存
	d.cache.Add(contentHash, true)

	// 添加到数据库
	return d.dbProvider.AddContentHash(contentHash, item.URL)
}

// Name 返回策略名称
func (d *ContentHashDeduplicator) Name() string {
	return "内容哈希去重"
}

// 生成内容哈希
func generateContentHash(content string) string {
	normalized := normalizeText(content)
	h := sha256.New()
	h.Write([]byte(normalized))
	return hex.EncodeToString(h.Sum(nil))
}

// ---------------- 标题去重策略 ----------------

// TitleDeduplicator 基于标题的去重策略
type TitleDeduplicator struct {
	cache      *lru.Cache[string, bool]
	dbProvider DBProvider
}

// NewTitleDeduplicator 创建标题去重器
func NewTitleDeduplicator(cacheSize int, dbProvider DBProvider) (*TitleDeduplicator, error) {
	cache, err := lru.New[string, bool](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("创建缓存失败: %v", err)
	}

	return &TitleDeduplicator{
		cache:      cache,
		dbProvider: dbProvider,
	}, nil
}

// IsDuplicate 检查标题是否已存在
func (d *TitleDeduplicator) IsDuplicate(item *Paper) (bool, float64, error) {
	normalizedTitle := normalizeText(item.Title)

	// 首先检查缓存
	if _, found := d.cache.Get(normalizedTitle); found {
		return true, 1.0, nil
	}

	// 查询数据库
	exists, err := d.dbProvider.TitleExists(normalizedTitle)
	if err != nil {
		return false, 0, err
	}

	if exists {
		// 添加到缓存
		d.cache.Add(normalizedTitle, true)
	}

	return exists, 1.0, nil
}

// AddItem 将标题添加到索引
func (d *TitleDeduplicator) AddItem(item *Paper) error {
	normalizedTitle := normalizeText(item.Title)

	// 添加到缓存
	d.cache.Add(normalizedTitle, true)

	// 添加到数据库
	return d.dbProvider.AddTitle(normalizedTitle, item.URL)
}

// Name 返回策略名称
func (d *TitleDeduplicator) Name() string {
	return "标题去重"
}

// ---------------- 相似度去重策略 ----------------

// SimilarityDeduplicator 基于内容相似度的去重策略
type SimilarityDeduplicator struct {
	threshold  float64
	cache      *lru.Cache[string, []string] // item_id -> tokens
	dbProvider DBProvider
}

// NewSimilarityDeduplicator 创建相似度去重器
func NewSimilarityDeduplicator(threshold float64, cacheSize int, dbProvider DBProvider) (*SimilarityDeduplicator, error) {
	cache, err := lru.New[string, []string](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("创建缓存失败: %v", err)
	}

	return &SimilarityDeduplicator{
		threshold:  threshold,
		cache:      cache,
		dbProvider: dbProvider,
	}, nil
}

// IsDuplicate 检查内容相似度
func (d *SimilarityDeduplicator) IsDuplicate(item *Paper) (bool, float64, error) {
	// 对内容进行分词
	itemTokens := tokenize(item.Description)

	// 首先在缓存中查找相似项
	highestSimilarity := 0.0

	// 查询最近的20条记录
	contentItems, err := d.dbProvider.GetLatestContentItems(20)
	if err != nil {
		return false, 0, err
	}

	for _, contentItem := range contentItems {
		// 计算相似度
		similarity := calculateJaccardSimilarity(itemTokens, contentItem.ContentTokens)

		if similarity > highestSimilarity {
			highestSimilarity = similarity
		}

		if similarity >= d.threshold {
			return true, similarity, nil
		}
	}

	return false, highestSimilarity, nil
}

// AddItem 将内容添加到索引
func (d *SimilarityDeduplicator) AddItem(item *Paper) error {
	// 对内容进行分词
	tokens := tokenize(item.Description)

	// 添加到缓存
	d.cache.Add(item.URL, tokens)

	// 添加到数据库
	return d.dbProvider.AddContentItem(item.URL, item.Description, strings.Join(tokens, ","))
}

// Name 返回策略名称
func (d *SimilarityDeduplicator) Name() string {
	return "内容相似度去重"
}

// ---------------- 辅助函数 ----------------

// 标准化文本(移除空白和标点，转为小写)
func normalizeText(text string) string {
	// 替换HTML标签
	text = cleanHtml(text)

	// 移除标点和特殊字符
	punctRegex := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	text = punctRegex.ReplaceAllString(text, " ")

	// 转为小写并去除多余空白
	text = strings.ToLower(text)
	text = strings.Join(strings.Fields(text), " ")

	return text
}

// 分词
func tokenize(text string) []string {
	normalized := normalizeText(text)
	return strings.Fields(normalized)
}

// 计算Jaccard相似度
func calculateJaccardSimilarity(tokens1, tokens2 []string) float64 {
	if len(tokens1) == 0 || len(tokens2) == 0 {
		return 0.0
	}

	// 创建集合
	set1 := make(map[string]bool)
	set2 := make(map[string]bool)

	for _, token := range tokens1 {
		set1[token] = true
	}

	for _, token := range tokens2 {
		set2[token] = true
	}

	// 计算交集大小
	intersection := 0
	for token := range set1 {
		if set2[token] {
			intersection++
		}
	}

	// 计算并集大小
	union := len(set1) + len(set2) - intersection

	// 计算Jaccard相似度
	return float64(intersection) / float64(union)
}
