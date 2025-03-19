package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// AISummarizer 接口定义了 AI 摘要功能
type AISummarizer interface {
	Summarize(content string) (string, error)
}

// 工厂函数创建 AI 摘要器
func NewAISummarizer(config *Config) (AISummarizer, error) {
	switch config.AIProvider {
	case "openai":
		if config.OpenAIAPIKey == "" {
			return nil, fmt.Errorf("OpenAI API Key not set in configuration")
		}
		return NewOpenAISummarizer(config.OpenAIAPIKey), nil
	case "azure":
		if config.AzureOpenAIAPIKey == "" || config.AzureOpenAIEndpoint == "" {
			return nil, fmt.Errorf("Azure OpenAI API Key or Endpoint not set in configuration")
		}
		return NewAzureOpenAISummarizer(config.AzureOpenAIAPIKey, config.AzureOpenAIEndpoint), nil
	case "deepseek":
		if config.DeepSeekAPIKey == "" {
			return nil, fmt.Errorf("DeepSeek API Key not set in configuration")
		}
		return NewDeepSeekAISummarizer(config.DeepSeekAPIKey), nil
	case "mock":
		// 用于测试的模拟实现
		return &MockAISummarizer{}, nil
	default:
		return nil, fmt.Errorf("unsupported AI provider: %s", config.AIProvider)
	}
}

// MockAISummarizer 是用于测试的 AI 摘要器模拟实现
type MockAISummarizer struct{}

// Summarize 返回模拟的摘要结果
func (s *MockAISummarizer) Summarize(content string) (string, error) {
	// 记录请求内容
	logInput(content)

	// 模拟响应
	mockResponse := "这是一篇关于AI Agent领域的论文，主要探讨了最新进展和研究方向。论文提出了一种新的方法来提高Agent的性能，并通过实验证明了其有效性。这项研究对未来的AI Agent开发具有重要指导意义。"

	// 记录响应内容
	logOutput(mockResponse)

	return mockResponse, nil
}

// OpenAI 摘要器实现
type OpenAISummarizer struct {
	apiKey string
}

func NewOpenAISummarizer(apiKey string) *OpenAISummarizer {
	return &OpenAISummarizer{apiKey: apiKey}
}

// OpenAI 请求结构
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
	MaxTokens   int             `json:"max_tokens"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAI 响应结构
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// Summarize 使用 OpenAI 生成摘要
func (s *OpenAISummarizer) Summarize(content string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	// 构建请求
	messages := []OpenAIMessage{
		{
			Role:    "system",
			Content: "你是一个有用的信息摘要AI助手，你可以从冗长的信息中摘要中找出关键点，并思考对未来的影响。你应该用中文回答我",
		},
		{
			Role:    "user",
			Content: content + " summary into short paragraph within 100 words",
		},
	}

	// 记录请求内容
	logInput(fmt.Sprintf("System: %s\nUser: %s",
		messages[0].Content,
		messages[1].Content))

	requestBody := OpenAIRequest{
		Model:       "gpt-3.5-turbo",
		Messages:    messages,
		Temperature: 0,
		MaxTokens:   300,
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
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	// 解析响应
	var response OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// 检查错误
	if response.Error.Message != "" {
		return "", fmt.Errorf("API error: %s", response.Error.Message)
	}

	// 检查是否有结果
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no summary generated")
	}

	result := response.Choices[0].Message.Content

	// 记录响应内容
	logOutput(result)

	return result, nil
}

// Azure OpenAI 实现
type AzureOpenAISummarizer struct {
	apiKey   string
	endpoint string
}

func NewAzureOpenAISummarizer(apiKey, endpoint string) *AzureOpenAISummarizer {
	return &AzureOpenAISummarizer{
		apiKey:   apiKey,
		endpoint: endpoint,
	}
}

// Azure OpenAI 实现 Summarize 方法
func (s *AzureOpenAISummarizer) Summarize(content string) (string, error) {
	// 记录请求内容
	logInput(content)

	// Azure OpenAI 调用逻辑与 OpenAI 类似，但 URL 和认证不同
	// 实际实现时可参考 Azure OpenAI 文档进行适配
	// 这里省略具体实现，返回占位符
	result := "Azure OpenAI 摘要 (未实现)"

	// 记录响应内容
	logOutput(result)

	return result, nil
}

// DeepSeek AI 摘要器实现
type DeepSeekAISummarizer struct {
	apiKey string
}

func NewDeepSeekAISummarizer(apiKey string) *DeepSeekAISummarizer {
	return &DeepSeekAISummarizer{apiKey: apiKey}
}

// DeepSeek 请求结构
type DeepSeekRequest struct {
	Model    string            `json:"model"`
	Messages []DeepSeekMessage `json:"messages"`
	Stream   bool              `json:"stream"`
}

type DeepSeekMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// DeepSeek 响应结构
type DeepSeekResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// Summarize 使用 DeepSeek 生成摘要
func (s *DeepSeekAISummarizer) Summarize(content string) (string, error) {
	url := "https://api.deepseek.com/chat/completions"

	// 构建请求
	messages := []DeepSeekMessage{
		{
			Role:    "system",
			Content: "你是一个有用的信息摘要AI助手，你可以从科学论文摘要中找出关键点。你应该用中文回答我",
		},
		{
			Role:    "user",
			Content: content + " summary into short paragraph within 100 words",
		},
	}

	// 记录请求内容
	logInput(fmt.Sprintf("System: %s\nUser: %s",
		messages[0].Content,
		messages[1].Content))

	requestBody := DeepSeekRequest{
		Model:    "deepseek-chat",
		Messages: messages,
		Stream:   false,
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
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	// 解析响应
	var response DeepSeekResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// 检查是否有结果
	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no summary generated")
	}

	result := response.Choices[0].Message.Content

	// 记录响应内容
	logOutput(result)

	return result, nil
}

// 日志辅助函数 - 记录输入内容
func logInput(content string) {
	// 清理HTML内容
	content = cleanHtml(content)

	// 截断过长的内容
	if len(content) > 1000 {
		content = content[:997] + "..."
	}

	// 日志分隔符
	log.Println("\n============ AI模型请求内容 ============")
	log.Println(content)
	log.Println("======================================\n")
}

// 日志辅助函数 - 记录输出内容
func logOutput(content string) {
	// 截断过长的内容
	if len(content) > 1000 {
		content = content[:997] + "..."
	}

	// 日志分隔符
	log.Println("\n============ AI模型响应内容 ============")
	log.Println(content)
	log.Println("======================================\n")
}

// 辅助函数 - 清理HTML内容
func cleanHtml(content string) string {
	// 基本HTML标签清理
	content = strings.ReplaceAll(content, "<br>", "\n")
	content = strings.ReplaceAll(content, "<br/>", "\n")
	content = strings.ReplaceAll(content, "<br />", "\n")

	// 尝试去除其他HTML标签
	tagRegex := regexp.MustCompile("<[^>]*>")
	content = tagRegex.ReplaceAllString(content, "")

	// HTML实体解码
	content = html.UnescapeString(content)

	return content
}
