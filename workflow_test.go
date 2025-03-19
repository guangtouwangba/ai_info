package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// MockRSSFetcher 是用于测试的 RSS 获取器模拟实现
type MockRSSFetcher struct{}

// Fetch 返回模拟的论文数据
func (f *MockRSSFetcher) Fetch(url string, maxResults int) ([]Paper, error) {
	// 返回一个模拟的论文数据
	papers := []Paper{
		{
			Title:       "测试论文：AI Agent 进展",
			URL:         "https://arxiv.org/abs/test.12345",
			Description: "这是一个关于 AI Agent 进展的测试论文摘要。包含了最新的研究方向和实验结果。",
			Summary:     "这是一个关于 AI Agent 进展的测试论文摘要。包含了最新的研究方向和实验结果。",
			SavedAt:     time.Now().Format(time.RFC3339),
		},
	}
	return papers, nil
}

// TestWorkflowWithMarkdown 测试工作流是否能正常执行并保存到 Markdown 文件
func TestWorkflowWithMarkdown(t *testing.T) {
	// 设置测试输出目录
	testOutputDir := "test_output"
	os.MkdirAll(testOutputDir, 0755)
	defer func() {
		// 可以选择在测试结束后清理目录
		// os.RemoveAll(testOutputDir)
	}()

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

	// 验证是否生成了 Markdown 文件
	files, err := os.ReadDir(testOutputDir)
	if err != nil {
		t.Fatalf("无法读取输出目录: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("未生成任何 Markdown 文件")
	}

	// 读取生成的文件内容进行验证
	filePath := filepath.Join(testOutputDir, files[0].Name())
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("无法读取生成的文件: %v", err)
	}

	// 验证文件内容是否包含预期的信息
	contentStr := string(content)
	expectedContents := []string{
		"# 测试论文：AI Agent 进展",
		"https://arxiv.org/abs/test.12345",
		"## 原始摘要",
		"## AI 摘要",
	}

	for _, expected := range expectedContents {
		if !contains(contentStr, expected) {
			t.Errorf("生成的文件内容缺少期望的文本: %s", expected)
		}
	}

	t.Logf("测试成功，生成的 Markdown 文件: %s", filePath)
}

// contains 检查字符串是否包含子串
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
