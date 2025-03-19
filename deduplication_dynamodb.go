package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBProvider 实现基于DynamoDB的数据库提供者
type DynamoDBProvider struct {
	client      *dynamodb.Client
	tablePrefix string
	region      string
	mu          sync.Mutex
}

// NewDynamoDBProvider 创建一个新的DynamoDB提供者
func NewDynamoDBProvider(tablePrefix string, config *DuplicationConfig) (*DynamoDBProvider, error) {
	region := config.AWSRegion
	if region == "" {
		// 尝试从主配置获取
		region = "us-east-1"
	}

	profile := config.AWSProfile
	if profile == "" {
		// 尝试从主配置获取
		profile = "default"
	}

	// 如果tablePrefix为空，使用一个默认值
	if tablePrefix == "" {
		tablePrefix = "dedup_"
	}

	log.Printf("初始化DynamoDB去重提供者，区域: %s, 表前缀: %s", region, tablePrefix)

	// 加载AWS配置
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(region),
		awsconfig.WithSharedConfigProfile(profile),
	)
	if err != nil {
		return nil, fmt.Errorf("无法加载AWS配置: %v", err)
	}

	// 创建DynamoDB客户端
	client := dynamodb.NewFromConfig(cfg)

	provider := &DynamoDBProvider{
		client:      client,
		tablePrefix: tablePrefix,
		region:      region,
	}

	if err := provider.Initialize(); err != nil {
		return nil, err
	}

	return provider, nil
}

// Initialize 初始化DynamoDB表
func (p *DynamoDBProvider) Initialize() error {
	tables := []struct {
		name       string
		hashKey    string
		attributes map[string]types.ScalarAttributeType
	}{
		{
			name:    p.tablePrefix + "url_index",
			hashKey: "url",
			attributes: map[string]types.ScalarAttributeType{
				"url": types.ScalarAttributeTypeS,
			},
		},
		{
			name:    p.tablePrefix + "content_hash_index",
			hashKey: "content_hash",
			attributes: map[string]types.ScalarAttributeType{
				"content_hash": types.ScalarAttributeTypeS,
			},
		},
		{
			name:    p.tablePrefix + "title_index",
			hashKey: "normalized_title",
			attributes: map[string]types.ScalarAttributeType{
				"normalized_title": types.ScalarAttributeTypeS,
			},
		},
		{
			name:    p.tablePrefix + "content_index",
			hashKey: "item_id",
			attributes: map[string]types.ScalarAttributeType{
				"item_id": types.ScalarAttributeTypeS,
			},
		},
	}

	for _, table := range tables {
		exist, err := p.tableExists(table.name)
		if err != nil {
			return fmt.Errorf("检查表 %s 是否存在时出错: %v", table.name, err)
		}

		if !exist {
			log.Printf("创建DynamoDB表: %s", table.name)
			if err := p.createTable(table.name, table.hashKey, table.attributes); err != nil {
				return fmt.Errorf("创建表 %s 失败: %v", table.name, err)
			}

			// 等待表变为活动状态
			if err := p.waitForTableActive(table.name); err != nil {
				return fmt.Errorf("等待表 %s 活动时出错: %v", table.name, err)
			}
			log.Printf("表 %s 创建成功", table.name)
		}
	}

	return nil
}

// tableExists 检查表是否存在
func (p *DynamoDBProvider) tableExists(tableName string) (bool, error) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	_, err := p.client.DescribeTable(context.TODO(), input)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// createTable 创建新表
func (p *DynamoDBProvider) createTable(tableName, hashKey string, attributes map[string]types.ScalarAttributeType) error {
	// 准备属性定义
	attrDefs := make([]types.AttributeDefinition, 0, len(attributes))
	for attrName, attrType := range attributes {
		attrDefs = append(attrDefs, types.AttributeDefinition{
			AttributeName: aws.String(attrName),
			AttributeType: attrType,
		})
	}

	// 准备键模式
	keySchema := []types.KeySchemaElement{
		{
			AttributeName: aws.String(hashKey),
			KeyType:       types.KeyTypeHash,
		},
	}

	// 创建表请求
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions:      attrDefs,
		KeySchema:                 keySchema,
		TableName:                 aws.String(tableName),
		BillingMode:               types.BillingModePayPerRequest, // 按需计费，自动扩展
		DeletionProtectionEnabled: aws.Bool(false),                // 允许删除表
	}

	_, err := p.client.CreateTable(context.TODO(), input)
	return err
}

// waitForTableActive 等待表变为活动状态
func (p *DynamoDBProvider) waitForTableActive(tableName string) error {
	// 等待表状态为ACTIVE
	waiter := dynamodb.NewTableExistsWaiter(p.client)
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}

	maxWaitTime := 5 * time.Minute
	log.Printf("等待表 %s 变为活动状态...", tableName)

	// 每5秒检查一次，最多等待5分钟
	err := waiter.Wait(context.TODO(), params, maxWaitTime)
	if err != nil {
		return err
	}

	return nil
}

// Close 关闭连接
func (p *DynamoDBProvider) Close() error {
	// DynamoDB客户端不需要显式关闭
	return nil
}

// URLExists 检查URL是否存在
func (p *DynamoDBProvider) URLExists(url string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "url_index"

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"url": &types.AttributeValueMemberS{Value: url},
		},
	}

	result, err := p.client.GetItem(context.TODO(), input)
	if err != nil {
		return false, fmt.Errorf("DynamoDB URL查询失败: %v", err)
	}

	return len(result.Item) > 0, nil
}

// AddURL 添加URL到索引
func (p *DynamoDBProvider) AddURL(url, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "url_index"

	item := struct {
		URL       string `dynamodbav:"url"`
		ItemID    string `dynamodbav:"item_id"`
		CreatedAt string `dynamodbav:"created_at"`
	}{
		URL:       url,
		ItemID:    itemID,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("DynamoDB编组URL项目失败: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = p.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("DynamoDB添加URL失败: %v", err)
	}

	return nil
}

// ContentHashExists 检查内容哈希是否存在
func (p *DynamoDBProvider) ContentHashExists(hash string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "content_hash_index"

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"content_hash": &types.AttributeValueMemberS{Value: hash},
		},
	}

	result, err := p.client.GetItem(context.TODO(), input)
	if err != nil {
		return false, fmt.Errorf("DynamoDB内容哈希查询失败: %v", err)
	}

	return len(result.Item) > 0, nil
}

// AddContentHash 添加内容哈希到索引
func (p *DynamoDBProvider) AddContentHash(hash, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "content_hash_index"

	item := struct {
		ContentHash string `dynamodbav:"content_hash"`
		ItemID      string `dynamodbav:"item_id"`
		CreatedAt   string `dynamodbav:"created_at"`
	}{
		ContentHash: hash,
		ItemID:      itemID,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("DynamoDB编组内容哈希项目失败: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = p.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("DynamoDB添加内容哈希失败: %v", err)
	}

	return nil
}

// TitleExists 检查标题是否存在
func (p *DynamoDBProvider) TitleExists(normalizedTitle string) (bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "title_index"

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"normalized_title": &types.AttributeValueMemberS{Value: normalizedTitle},
		},
	}

	result, err := p.client.GetItem(context.TODO(), input)
	if err != nil {
		return false, fmt.Errorf("DynamoDB标题查询失败: %v", err)
	}

	return len(result.Item) > 0, nil
}

// AddTitle 添加标题到索引
func (p *DynamoDBProvider) AddTitle(normalizedTitle, itemID string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "title_index"

	item := struct {
		NormalizedTitle string `dynamodbav:"normalized_title"`
		ItemID          string `dynamodbav:"item_id"`
		CreatedAt       string `dynamodbav:"created_at"`
	}{
		NormalizedTitle: normalizedTitle,
		ItemID:          itemID,
		CreatedAt:       time.Now().Format(time.RFC3339),
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("DynamoDB编组标题项目失败: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = p.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("DynamoDB添加标题失败: %v", err)
	}

	return nil
}

// GetLatestContentItems 获取最新的内容项目
func (p *DynamoDBProvider) GetLatestContentItems(limit int) ([]ContentItem, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "content_index"

	// 扫描表获取最新的项目
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		Limit:     aws.Int32(int32(limit)),
	}

	result, err := p.client.Scan(context.TODO(), scanInput)
	if err != nil {
		return nil, fmt.Errorf("DynamoDB获取内容项目失败: %v", err)
	}

	var items []ContentItem
	for _, item := range result.Items {
		var dynamoItem struct {
			ItemID        string `dynamodbav:"item_id"`
			Content       string `dynamodbav:"content"`
			ContentTokens string `dynamodbav:"content_tokens"`
		}

		if err := attributevalue.UnmarshalMap(item, &dynamoItem); err != nil {
			log.Printf("解组内容项目失败: %v", err)
			continue
		}

		contentItem := ContentItem{
			ItemID:  dynamoItem.ItemID,
			Content: dynamoItem.Content,
		}

		// 解析tokens
		if dynamoItem.ContentTokens != "" {
			contentItem.ContentTokens = strings.Split(dynamoItem.ContentTokens, ",")
		}

		items = append(items, contentItem)
	}

	return items, nil
}

// AddContentItem 添加内容项目到索引
func (p *DynamoDBProvider) AddContentItem(itemID, content, contentTokens string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	tableName := p.tablePrefix + "content_index"

	item := struct {
		ItemID        string `dynamodbav:"item_id"`
		Content       string `dynamodbav:"content"`
		ContentTokens string `dynamodbav:"content_tokens"`
		CreatedAt     string `dynamodbav:"created_at"`
	}{
		ItemID:        itemID,
		Content:       content,
		ContentTokens: contentTokens,
		CreatedAt:     time.Now().Format(time.RFC3339),
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("DynamoDB编组内容项目失败: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = p.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("DynamoDB添加内容项目失败: %v", err)
	}

	return nil
}
