package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestMainWithMarkdown 测试整个程序使用 Markdown 存储的端到端流程
func TestMainWithMarkdown(t *testing.T) {
	// 设置测试环境变量
	testOutputDir := "test_output_main"
	os.Setenv("OUTPUT_DIR", testOutputDir)
	//defer os.RemoveAll(testOutputDir) // 测试结束后清理

	// 保存当前命令行参数并在函数返回时恢复
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// 设置命令行参数用于测试
	os.Args = []string{
		"arxiv-agent",
		"-ai", "mock",
		"-storage", "markdown",
		"-max", "1",
	}

	// 通过手动执行相关测试步骤进行测试
	// 创建 RSS 获取器 (使用 mock 测试数据)
	rssFetcher := &MockRSSFetcher{}

	// 从 RSS 获取测试数据
	papers, err := rssFetcher.Fetch("test-url", 1)
	if err != nil {
		t.Fatalf("获取 RSS 数据失败: %v", err)
	}

	// 创建 Markdown 存储提供者
	storage := NewMarkdownStorageProvider(testOutputDir)

	// 创建 Mock AI 摘要器
	aiSummarizer := &MockAISummarizer{}

	// 处理测试论文
	for _, paper := range papers {
		// 生成摘要
		summary, err := aiSummarizer.Summarize(paper.Description)
		if err != nil {
			t.Fatalf("生成摘要失败: %v", err)
		}

		paper.AISummary = summary
		paper.SavedAt = time.Now().Format(time.RFC3339)

		// 保存论文
		if err := storage.Save(paper); err != nil {
			t.Fatalf("保存论文失败: %v", err)
		}
	}

	// 验证结果
	validateOutputFiles(t, testOutputDir)
}

// TestMarkdownStorage 专门测试 Markdown 存储功能
func TestMarkdownStorage(t *testing.T) {
	// 设置测试输出目录
	testOutputDir := "test_output_storage"
	defer os.RemoveAll(testOutputDir) // 测试结束后清理

	// 创建存储提供者
	storage := NewMarkdownStorageProvider(testOutputDir)

	// 创建测试论文
	paper := Paper{
		Title:       "测试论文标题",
		URL:         "https://example.com/test",
		Description: "这是一个测试描述。",
		Summary:     "这是原始摘要。",
		AISummary:   "这是 AI 生成的摘要。",
		SavedAt:     time.Now().Format(time.RFC3339),
	}

	// 保存论文
	err := storage.Save(paper)
	if err != nil {
		t.Fatalf("保存论文失败: %v", err)
	}

	// 验证结果
	validateOutputFiles(t, testOutputDir)
}

// TestDeepSeekWithMarkdown 测试使用 DeepSeek 和 Markdown 存储的端到端流程
func TestDeepSeekWithMarkdown(t *testing.T) {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("无法加载配置: %v", err)
	}

	// 检查是否设置了 DeepSeek API 密钥，如果没有则跳过测试
	if config.DeepSeekAPIKey == "" {
		t.Skip("跳过测试：未设置 DeepSeek API 密钥")
	}

	// 设置测试输出目录
	testOutputDir := "test_output_deepseek"
	os.Setenv("OUTPUT_DIR", testOutputDir)
	defer os.RemoveAll(testOutputDir) // 测试结束后清理

	// 通过手动执行相关测试步骤进行测试
	// 创建 RSS 获取器 (使用 mock 测试数据)
	rssFetcher := &MockRSSFetcher{}

	// 从 RSS 获取测试数据，使用测试URL
	testURL := "test-url"
	papers, err := rssFetcher.Fetch(testURL, 1)
	if err != nil {
		t.Fatalf("获取 RSS 数据失败: %v", err)
	}

	// 创建 Markdown 存储提供者
	storage := NewMarkdownStorageProvider(testOutputDir)

	// 创建 DeepSeek AI 摘要器
	aiSummarizer := NewDeepSeekAISummarizer(config.DeepSeekAPIKey)

	// 处理测试论文
	for _, paper := range papers {
		// 生成摘要
		summary, err := aiSummarizer.Summarize(paper.Description)
		if err != nil {
			t.Fatalf("生成摘要失败: %v", err)
		}

		paper.AISummary = summary
		paper.SavedAt = time.Now().Format(time.RFC3339)

		// 保存论文
		if err := storage.Save(paper); err != nil {
			t.Fatalf("保存论文失败: %v", err)
		}
	}

	// 验证结果
	validateOutputFiles(t, testOutputDir)
}

// validateOutputFiles 验证输出目录中是否存在预期的文件
func validateOutputFiles(t *testing.T, outputDir string) {
	// 验证目录是否存在
	_, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		t.Fatalf("输出目录不存在: %s", outputDir)
	}

	// 验证是否生成了文件
	files, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("无法读取输出目录: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("未生成任何文件")
	}

	// 验证文件扩展名
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			t.Errorf("生成的文件不是 Markdown 文件: %s", file.Name())
		}
	}

	// 验证文件内容
	filePath := filepath.Join(outputDir, files[0].Name())
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("无法读取生成的文件: %v", err)
	}

	contentStr := string(content)
	t.Logf("生成的文件内容: %s", contentStr)

	// 验证必要的标题和部分
	requiredSections := []string{
		"# ",       // 标题
		"**URL**",  // URL 部分
		"## 原始摘要",  // 原始摘要部分
		"## AI 摘要", // AI 摘要部分
		"## 元数据",   // 元数据部分
	}

	for _, section := range requiredSections {
		if !strings.Contains(contentStr, section) {
			t.Errorf("文件缺少必要的部分: %s", section)
		}
	}
}
