package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// StorageProvider 接口定义了保存数据的方法
type StorageProvider interface {
	Save(paper Paper) error
}

// 工厂函数创建存储提供者
func NewStorageProvider(config *Config) (StorageProvider, error) {
	switch config.StorageProvider {
	case "notion":
		if config.NotionAPIKey == "" || config.NotionDatabaseID == "" {
			return nil, fmt.Errorf("Notion API Key or Database ID not set in configuration")
		}
		return NewNotionStorageProvider(config.NotionAPIKey, config.NotionDatabaseID), nil
	case "markdown":
		outputDir := config.OutputDir
		if outputDir == "" {
			outputDir = "papers" // 默认输出目录
		}
		return NewMarkdownStorageProvider(outputDir), nil
	case "dynamodb":
		return NewDynamoDBStorageProvider(config)
	default:
		return nil, fmt.Errorf("unsupported storage provider: %s", config.StorageProvider)
	}
}

// Notion 存储提供者
type NotionStorageProvider struct {
	apiKey     string
	databaseID string
}

func NewNotionStorageProvider(apiKey, databaseID string) *NotionStorageProvider {
	return &NotionStorageProvider{
		apiKey:     apiKey,
		databaseID: databaseID,
	}
}

// Save 将论文保存到 Notion 数据库
func (p *NotionStorageProvider) Save(paper Paper) error {
	// 创建 Notion 页面
	pageID, err := p.createPage(paper)
	if err != nil {
		return fmt.Errorf("failed to create Notion page: %v", err)
	}

	// 向页面添加内容
	if err := p.appendContent(pageID, paper); err != nil {
		return fmt.Errorf("failed to append content to Notion page: %v", err)
	}

	return nil
}

// createPage 在 Notion 创建新页面
func (p *NotionStorageProvider) createPage(paper Paper) (string, error) {
	url := "https://api.notion.com/v1/pages"

	// 构建请求体，对应 Notion API 结构
	requestBody := map[string]interface{}{
		"parent": map[string]string{
			"database_id": p.databaseID,
		},
		"properties": map[string]interface{}{
			"title": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]string{
							"content": paper.Title,
						},
					},
				},
			},
			"description": map[string]interface{}{
				"rich_text": []map[string]interface{}{
					{
						"text": map[string]string{
							"content": paper.Description,
						},
					},
				},
			},
			"URL": map[string]string{
				"url": paper.URL,
			},
			"summary": map[string]interface{}{
				"rich_text": []map[string]interface{}{
					{
						"text": map[string]string{
							"content": paper.AISummary,
						},
					},
				},
			},
			"tag": map[string]interface{}{
				"multi_select": []map[string]string{
					{
						"name": "paper",
					},
				},
			},
			"comes from": map[string]interface{}{
				"multi_select": []map[string]string{
					{
						"name": "make",
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// 检查是否有错误
	if resp.StatusCode != http.StatusOK {
		errorMsg, _ := json.Marshal(response)
		return "", fmt.Errorf("API error: %s", string(errorMsg))
	}

	// 获取页面 ID
	pageID, ok := response["id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get page ID from response")
	}

	return pageID, nil
}

// appendContent 向 Notion 页面添加内容
func (p *NotionStorageProvider) appendContent(pageID string, paper Paper) error {
	url := "https://api.notion.com/v1/blocks/" + pageID + "/children"

	// 构建请求体，添加标题和摘要
	requestBody := map[string]interface{}{
		"children": []map[string]interface{}{
			{
				"object": "block",
				"type":   "heading_1",
				"heading_1": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]string{
								"content": paper.Title,
							},
						},
					},
				},
			},
			{
				"object": "block",
				"type":   "paragraph",
				"paragraph": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]string{
								"content": paper.AISummary,
							},
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)
	req.Header.Set("Notion-Version", "2022-06-28")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查是否有错误
	if resp.StatusCode != http.StatusOK {
		var response map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&response)
		errorMsg, _ := json.Marshal(response)
		return fmt.Errorf("API error: %s", string(errorMsg))
	}

	return nil
}

// Markdown 存储提供者
type MarkdownStorageProvider struct {
	outputDir string
}

func NewMarkdownStorageProvider(outputDir string) *MarkdownStorageProvider {
	return &MarkdownStorageProvider{
		outputDir: outputDir,
	}
}

// Save 将论文保存为 Markdown 文件
func (p *MarkdownStorageProvider) Save(paper Paper) error {
	// 确定保存日期，如果Paper中没有日期则使用当前时间
	saveTime := time.Now()
	if paper.SavedAt != "" {
		// 尝试解析已有的SavedAt字段
		parsedTime, err := time.Parse(time.RFC3339, paper.SavedAt)
		if err == nil {
			saveTime = parsedTime
		}
	}

	// 格式化日期作为目录名
	dateDir := saveTime.Format("2006-01-02")

	// 完整的目录路径
	fullDirPath := filepath.Join(p.outputDir, dateDir)

	// 确保输出目录存在
	if err := os.MkdirAll(fullDirPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// 生成文件名
	fileName := fmt.Sprintf("%s.md", sanitizeFileName(paper.Title))
	filePath := filepath.Join(fullDirPath, fileName)

	// 创建 Markdown 内容
	content := fmt.Sprintf("# %s\n\n", paper.Title)
	content += fmt.Sprintf("**URL**: %s\n\n", paper.URL)
	content += fmt.Sprintf("## 原始摘要\n\n%s\n\n", paper.Description)
	content += fmt.Sprintf("## AI 摘要\n\n%s\n\n", paper.AISummary)
	content += fmt.Sprintf("## 元数据\n\n")
	content += fmt.Sprintf("- **来源**: ArXiv\n")
	content += fmt.Sprintf("- **类型**: 论文\n")
	content += fmt.Sprintf("- **保存时间**: %s\n", paper.SavedAt)
	content += fmt.Sprintf("- **目录日期**: %s\n", dateDir)

	// 写入文件
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	log.Printf("已保存论文到 %s", filePath)

	// 生成或更新索引
	if err := p.GenerateIndex(); err != nil {
		log.Printf("警告: 生成索引失败: %v", err)
	}

	return nil
}

// 添加新函数sanitizeTextForMarkdown来清理将要写入markdown中的文本
func sanitizeTextForMarkdown(text string) string {
	// 确保字符是有效的UTF-8字符
	if !utf8.ValidString(text) {
		var validContent strings.Builder
		for _, r := range text {
			if r != utf8.RuneError {
				validContent.WriteRune(r)
			}
		}
		text = validContent.String()
	}

	// 替换可能导致Markdown解析问题的字符
	text = strings.ReplaceAll(text, "|", "\\|") // 转义竖线，防止表格混淆
	text = strings.ReplaceAll(text, "*", "\\*") // 转义星号，防止强调符号混淆
	text = strings.ReplaceAll(text, "_", "\\_") // 转义下划线，防止强调符号混淆
	text = strings.ReplaceAll(text, "`", "\\`") // 转义反引号，防止代码块混淆

	// 移除所有控制字符和不可见字符(除了空格)
	var cleanContent strings.Builder
	for _, r := range text {
		if !unicode.IsControl(r) || r == ' ' || r == '\t' || r == '\n' {
			cleanContent.WriteRune(r)
		}
	}

	return cleanContent.String()
}

// GenerateIndex 生成索引文件，列出所有日期目录和论文
func (p *MarkdownStorageProvider) GenerateIndex() error {
	// 确保输出目录存在
	if err := os.MkdirAll(p.outputDir, 0755); err != nil {
		return fmt.Errorf("创建输出目录失败: %v", err)
	}

	// 索引文件路径
	indexPath := filepath.Join(p.outputDir, "README.md")

	// 收集所有日期目录
	dirs, err := os.ReadDir(p.outputDir)
	if err != nil {
		return fmt.Errorf("读取目录失败: %v", err)
	}

	// 按日期排序的目录列表（最新的在前）
	var dateDirs []string
	for _, dir := range dirs {
		if dir.IsDir() && isDirDateFormat(dir.Name()) {
			dateDirs = append(dateDirs, dir.Name())
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dateDirs)))

	// 生成索引内容
	var content strings.Builder

	// 写入带有UTF-8 BOM的标记，使某些编辑器能够正确识别UTF-8编码
	// 注意：通常这不是必需的，因为UTF-8不需要BOM，但对于某些特定软件可能有帮助
	// content.WriteString("\uFEFF")

	content.WriteString("# 最新信息\n\n")
	content.WriteString("按日期组织的内容集合，最新的日期在前。每个条目包含标题、摘要、发布时间和原始链接。\n\n")

	if len(dateDirs) == 0 {
		content.WriteString("*暂无内容*\n")
	} else {
		// 最新日期的内容
		if len(dateDirs) > 0 {
			latestDir := dateDirs[0]
			content.WriteString(fmt.Sprintf("## 最新内容 (%s)\n\n", latestDir))

			papers, err := listPapersInDir(filepath.Join(p.outputDir, latestDir))
			if err == nil && len(papers) > 0 {
				for _, paper := range papers {
					relativePath := filepath.Join(latestDir, paper.filename)

					// 使用sanitizeTextForMarkdown确保标题不含任何可能导致乱码的字符
					sanitizedTitle := sanitizeTextForMarkdown(paper.title)

					// 增强展示，包含标题、摘要、时间和链接
					content.WriteString(fmt.Sprintf("### [%s](%s)\n\n", sanitizedTitle, relativePath))

					if paper.summary != "" {
						// 确保摘要文本也经过清理
						sanitizedSummary := sanitizeTextForMarkdown(paper.summary)
						content.WriteString(fmt.Sprintf("**摘要**: %s\n\n", sanitizedSummary))
					}

					if paper.url != "" {
						content.WriteString(fmt.Sprintf("**原始链接**: [查看原文](%s)\n", paper.url))
					}

					if paper.savedAt != "" {
						content.WriteString(fmt.Sprintf("**收录时间**: %s\n", paper.savedAt))
					}

					content.WriteString("\n---\n\n")
				}
			} else {
				content.WriteString("*该日期下暂无内容*\n")
			}
			content.WriteString("\n")
		}

		// 按日期列出所有内容
		content.WriteString("## 历史内容\n\n")
		for _, dir := range dateDirs {
			// 跳过最新日期（因为已经在"最新内容"部分列出）
			if dir == dateDirs[0] {
				continue
			}

			papers, err := listPapersInDir(filepath.Join(p.outputDir, dir))
			if err != nil {
				continue
			}

			content.WriteString(fmt.Sprintf("### %s (%d篇)\n\n", dir, len(papers)))

			for _, paper := range papers {
				relativePath := filepath.Join(dir, paper.filename)

				// 确保标题经过清理
				sanitizedTitle := sanitizeTextForMarkdown(paper.title)

				// 简洁展示历史内容，但仍包含关键信息
				content.WriteString(fmt.Sprintf("- **[%s](%s)** ", sanitizedTitle, relativePath))

				if paper.url != "" {
					content.WriteString(fmt.Sprintf("| [原文](%s) ", paper.url))
				}

				if paper.savedAt != "" {
					content.WriteString(fmt.Sprintf("| 时间: %s ", paper.savedAt))
				}

				content.WriteString("\n")

				if paper.summary != "" {
					// 确保摘要文本也经过清理
					sanitizedSummary := sanitizeTextForMarkdown(paper.summary)
					content.WriteString(fmt.Sprintf("  %s\n", sanitizedSummary))
				}

				content.WriteString("\n")
			}
		}
	}

	// 写入索引文件，确保用UTF-8编码
	if err := os.WriteFile(indexPath, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("写入索引文件失败: %v", err)
	}

	log.Printf("已更新索引文件: %s", indexPath)
	return nil
}

// 辅助结构体和函数

// 判断目录名是否为日期格式 (YYYY-MM-DD)
func isDirDateFormat(name string) bool {
	_, err := time.Parse("2006-01-02", name)
	return err == nil
}

// 论文信息，扩展结构以包含更多元数据
type paperInfo struct {
	title    string
	filename string
	summary  string // AI摘要
	url      string // 原始链接
	savedAt  string // 保存时间
}

// 列出目录中的所有论文，并获取详细信息
func listPapersInDir(dirPath string) ([]paperInfo, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var papers []paperInfo
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			// 初始化论文信息
			info := paperInfo{
				title:    strings.TrimSuffix(file.Name(), ".md"),
				filename: file.Name(),
				summary:  "",
				url:      "",
				savedAt:  "",
			}

			filePath := filepath.Join(dirPath, file.Name())
			// 使用增强的文件读取函数
			content, err := readFileWithCorrectEncoding(filePath)
			if err == nil {
				// 解析文件内容以提取元数据
				lines := strings.Split(content, "\n")

				// 提取标题（第一行，格式为 # 标题）
				if len(lines) > 0 && strings.HasPrefix(lines[0], "# ") {
					info.title = strings.TrimPrefix(lines[0], "# ")
				}

				// 创建两个变量分别存储原始摘要和AI摘要
				var originalSummary, aiSummary string

				// 提取URL（通常在第三行，格式为 **URL**: http://...）
				for i, line := range lines {
					if strings.HasPrefix(line, "**URL**: ") {
						info.url = strings.TrimPrefix(line, "**URL**: ")
					}

					// 查找原始摘要部分
					if line == "## 原始摘要" && i+1 < len(lines) {
						var summaryBuilder strings.Builder
						summaryLine := i + 1

						// 收集摘要内容直到遇到空行或新标题
						lineCount := 0
						for ; summaryLine < len(lines) && lineCount < 5; summaryLine++ {
							currentLine := lines[summaryLine]
							// 跳过空行
							if currentLine == "" {
								continue
							}
							// 如果遇到新的部分标题，结束收集
							if strings.HasPrefix(currentLine, "##") {
								break
							}

							if summaryBuilder.Len() > 0 {
								summaryBuilder.WriteString(" ") // 添加空格连接不同行
							}
							summaryBuilder.WriteString(currentLine)
							lineCount++
						}

						originalSummary = summaryBuilder.String()
					}

					// 查找AI摘要部分
					if line == "## AI 摘要" && i+1 < len(lines) {
						var summaryBuilder strings.Builder
						summaryLine := i + 1

						// 收集AI摘要内容
						lineCount := 0
						for ; summaryLine < len(lines) && lineCount < 5; summaryLine++ {
							currentLine := lines[summaryLine]
							// 跳过空行
							if currentLine == "" {
								continue
							}
							// 如果遇到新的部分标题，结束收集
							if strings.HasPrefix(currentLine, "##") {
								break
							}

							if summaryBuilder.Len() > 0 {
								summaryBuilder.WriteString(" ") // 添加空格连接不同行
							}
							summaryBuilder.WriteString(currentLine)
							lineCount++
						}

						aiSummary = summaryBuilder.String()
					}

					// 查找保存时间
					if strings.HasPrefix(line, "- **保存时间**: ") {
						info.savedAt = strings.TrimPrefix(line, "- **保存时间**: ")
						// 尝试格式化时间为更友好的显示
						if t, err := time.Parse(time.RFC3339, info.savedAt); err == nil {
							info.savedAt = t.Format("2006-01-02 15:04")
						}
					}
				}

				// 选择更有信息量的摘要
				chosenSummary := aiSummary

				// 如果AI摘要为空或包含模板文本"主要探讨了最新进展和研究方向"，则使用原始摘要
				if aiSummary == "" || strings.Contains(aiSummary, "主要探讨了最新进展和研究方向") {
					if originalSummary != "" {
						chosenSummary = originalSummary
					}
				}

				// 处理HTML内容，提取纯文本
				if strings.Contains(chosenSummary, "<") && strings.Contains(chosenSummary, ">") {
					// 简单的HTML标签清理
					// 1. 移除所有<开头到>结尾的内容
					htmlTagRegex := regexp.MustCompile("<[^>]*>")
					chosenSummary = htmlTagRegex.ReplaceAllString(chosenSummary, "")

					// 2. 处理HTML实体
					chosenSummary = strings.ReplaceAll(chosenSummary, "&nbsp;", " ")
					chosenSummary = strings.ReplaceAll(chosenSummary, "&lt;", "<")
					chosenSummary = strings.ReplaceAll(chosenSummary, "&gt;", ">")
					chosenSummary = strings.ReplaceAll(chosenSummary, "&amp;", "&")
					chosenSummary = strings.ReplaceAll(chosenSummary, "&quot;", "\"")

					// 3. 移除多余空格
					spaceRegex := regexp.MustCompile(`\s+`)
					chosenSummary = spaceRegex.ReplaceAllString(chosenSummary, " ")
					chosenSummary = strings.TrimSpace(chosenSummary)
				}

				// 处理摘要长度，保证是完整的中文或英文句子
				if len(chosenSummary) > 150 {
					// 尝试找一个合适的断句点（句号、问号、感叹号等）
					cutoffIndex := 150
					for j := cutoffIndex; j >= cutoffIndex-30 && j >= 0; j-- {
						if j >= len(chosenSummary) {
							continue
						}
						char := rune(chosenSummary[j])
						if char == '.' || char == '。' || char == '!' || char == '！' ||
							char == '?' || char == '？' || char == ';' || char == '；' ||
							char == ',' || char == '，' {
							cutoffIndex = j + 1
							break
						}
					}
					// 确保cutoffIndex不会超出范围
					if cutoffIndex > len(chosenSummary) {
						cutoffIndex = len(chosenSummary)
					}
					info.summary = chosenSummary[:cutoffIndex] + "..."
				} else {
					info.summary = chosenSummary
				}
			}

			papers = append(papers, info)
		}
	}

	return papers, nil
}

// readFileWithCorrectEncoding 读取文件内容并确保正确处理UTF-8编码
func readFileWithCorrectEncoding(filePath string) (string, error) {
	rawBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// 检测并处理UTF-8 BOM（字节顺序标记）
	if len(rawBytes) >= 3 && rawBytes[0] == 0xEF && rawBytes[1] == 0xBB && rawBytes[2] == 0xBF {
		// 如果文件以UTF-8 BOM开头，则跳过这3个字节
		rawBytes = rawBytes[3:]
	}

	// 尝试将内容解码为UTF-8
	content := string(rawBytes)

	// 检查是否包含无效的UTF-8序列
	if !utf8.ValidString(content) {
		// 如果不是有效的UTF-8，尝试其他编码处理
		// 这里我们使用一个更强大的方法来处理无效UTF-8
		var validContent strings.Builder
		for i := 0; i < len(content); {
			r, size := utf8.DecodeRuneInString(content[i:])
			if r == utf8.RuneError && size == 1 {
				// 处理无效字符：忽略或替换为其他字符
				// 这里选择忽略
				i++
				continue
			}
			validContent.WriteRune(r)
			i += size
		}
		content = validContent.String()
	}

	// 移除任何控制字符（除常见格式控制外）
	var cleanContent strings.Builder
	for _, r := range content {
		// 保留常见格式控制字符（空格、制表符、换行符等）
		if !unicode.IsControl(r) || r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			cleanContent.WriteRune(r)
		}
	}

	return cleanContent.String(), nil
}

// sanitizeFileName 清理文件名，替换不合法字符
func sanitizeFileName(name string) string {
	// 这里简化处理，实际中可能需要更复杂的清理逻辑
	replacer := strings.NewReplacer(
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
		" ", "_",
	)
	return replacer.Replace(name)
}

// DynamoDBStorageProvider 实现基于DynamoDB的存储提供者
type DynamoDBStorageProvider struct {
	client    *dynamodb.Client
	tableName string
	region    string
}

// NewDynamoDBStorageProvider 创建一个新的DynamoDB存储提供者
func NewDynamoDBStorageProvider(config *Config) (StorageProvider, error) {
	region := config.AWSRegion
	if region == "" {
		region = "us-east-1"
	}

	profile := config.AWSProfile
	if profile == "" {
		profile = "default"
	}

	tableName := config.AWSDynamoDBTable
	if tableName == "" {
		tableName = "arxiv_papers"
	}

	log.Printf("使用DynamoDB作为存储提供者，区域: %s, 配置文件: %s, 表名: %s", region, profile, tableName)

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

	provider := &DynamoDBStorageProvider{
		client:    client,
		tableName: tableName,
		region:    region,
	}

	// 确保表存在
	if err := provider.ensureTableExists(); err != nil {
		return nil, err
	}

	return provider, nil
}

// ensureTableExists 确保存储论文的表已存在
func (p *DynamoDBStorageProvider) ensureTableExists() error {
	// 检查表是否存在
	exists, err := p.tableExists(p.tableName)
	if err != nil {
		return fmt.Errorf("检查表是否存在时出错: %v", err)
	}

	if !exists {
		log.Printf("创建DynamoDB表: %s", p.tableName)

		// 创建表
		input := &dynamodb.CreateTableInput{
			TableName: aws.String(p.tableName),
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("url"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("saved_at"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("url"),
					KeyType:       types.KeyTypeHash,
				},
			},
			GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
				{
					IndexName: aws.String("saved_at-index"),
					KeySchema: []types.KeySchemaElement{
						{
							AttributeName: aws.String("saved_at"),
							KeyType:       types.KeyTypeHash,
						},
					},
					Projection: &types.Projection{
						ProjectionType: types.ProjectionTypeAll,
					},
					ProvisionedThroughput: &types.ProvisionedThroughput{
						ReadCapacityUnits:  aws.Int64(5),
						WriteCapacityUnits: aws.Int64(5),
					},
				},
			},
			BillingMode: types.BillingModeProvisioned,
			ProvisionedThroughput: &types.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
		}

		_, err = p.client.CreateTable(context.TODO(), input)
		if err != nil {
			return fmt.Errorf("创建表失败: %v", err)
		}

		// 等待表变为活动状态
		if err := p.waitForTableActive(p.tableName); err != nil {
			return fmt.Errorf("等待表活动时出错: %v", err)
		}

		log.Printf("DynamoDB表创建成功: %s", p.tableName)
	}

	return nil
}

// tableExists 检查表是否存在
func (p *DynamoDBStorageProvider) tableExists(tableName string) (bool, error) {
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

// waitForTableActive 等待表变为活动状态
func (p *DynamoDBStorageProvider) waitForTableActive(tableName string) error {
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

// Save 将论文保存到DynamoDB
func (p *DynamoDBStorageProvider) Save(paper Paper) error {
	// 为DynamoDB准备Paper结构
	type DynamoPaper struct {
		URL         string `dynamodbav:"url"`
		Title       string `dynamodbav:"title"`
		Description string `dynamodbav:"description"`
		Summary     string `dynamodbav:"summary"`
		AISummary   string `dynamodbav:"ai_summary"`
		SavedAt     string `dynamodbav:"saved_at"`
	}

	// 确保SavedAt有值
	savedAt := paper.SavedAt
	if savedAt == "" {
		savedAt = time.Now().Format(time.RFC3339)
	}

	// 准备要保存的数据
	dynamoPaper := DynamoPaper{
		URL:         paper.URL,
		Title:       paper.Title,
		Description: paper.Description,
		Summary:     paper.Summary,
		AISummary:   paper.AISummary,
		SavedAt:     savedAt,
	}

	// 转换为DynamoDB属性值映射
	av, err := attributevalue.MarshalMap(dynamoPaper)
	if err != nil {
		return fmt.Errorf("编组Paper项目失败: %v", err)
	}

	// 构建PutItem请求
	input := &dynamodb.PutItemInput{
		TableName: aws.String(p.tableName),
		Item:      av,
	}

	// 执行PutItem操作
	_, err = p.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("保存到DynamoDB失败: %v", err)
	}

	log.Printf("已将论文 [%s] 保存到DynamoDB表 %s", paper.Title, p.tableName)
	return nil
}

// GetPaperByURL 通过URL获取论文(可选的辅助方法)
func (p *DynamoDBStorageProvider) GetPaperByURL(url string) (*Paper, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(p.tableName),
		Key: map[string]types.AttributeValue{
			"url": &types.AttributeValueMemberS{Value: url},
		},
	}

	result, err := p.client.GetItem(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("从DynamoDB获取论文失败: %v", err)
	}

	if len(result.Item) == 0 {
		return nil, nil // 未找到论文
	}

	// 解析结果
	var dynamoPaper struct {
		URL         string `dynamodbav:"url"`
		Title       string `dynamodbav:"title"`
		Description string `dynamodbav:"description"`
		Summary     string `dynamodbav:"summary"`
		AISummary   string `dynamodbav:"ai_summary"`
		SavedAt     string `dynamodbav:"saved_at"`
	}

	if err := attributevalue.UnmarshalMap(result.Item, &dynamoPaper); err != nil {
		return nil, fmt.Errorf("解组Paper项目失败: %v", err)
	}

	// 转换为Paper结构
	paper := &Paper{
		URL:         dynamoPaper.URL,
		Title:       dynamoPaper.Title,
		Description: dynamoPaper.Description,
		Summary:     dynamoPaper.Summary,
		AISummary:   dynamoPaper.AISummary,
		SavedAt:     dynamoPaper.SavedAt,
	}

	return paper, nil
}

// ListRecentPapers 列出最近保存的论文(可选的辅助方法)
func (p *DynamoDBStorageProvider) ListRecentPapers(limit int) ([]Paper, error) {
	if limit <= 0 {
		limit = 10 // 默认限制
	}

	// 构建查询
	input := &dynamodb.ScanInput{
		TableName:      aws.String(p.tableName),
		Limit:          aws.Int32(int32(limit)),
		ConsistentRead: aws.Bool(false),
	}

	result, err := p.client.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("扫描DynamoDB失败: %v", err)
	}

	var papers []Paper
	for _, item := range result.Items {
		var dynamoPaper struct {
			URL         string `dynamodbav:"url"`
			Title       string `dynamodbav:"title"`
			Description string `dynamodbav:"description"`
			Summary     string `dynamodbav:"summary"`
			AISummary   string `dynamodbav:"ai_summary"`
			SavedAt     string `dynamodbav:"saved_at"`
		}

		if err := attributevalue.UnmarshalMap(item, &dynamoPaper); err != nil {
			log.Printf("解组Paper项目失败: %v", err)
			continue
		}

		paper := Paper{
			URL:         dynamoPaper.URL,
			Title:       dynamoPaper.Title,
			Description: dynamoPaper.Description,
			Summary:     dynamoPaper.Summary,
			AISummary:   dynamoPaper.AISummary,
			SavedAt:     dynamoPaper.SavedAt,
		}

		papers = append(papers, paper)
	}

	return papers, nil
}
