package main

import (
	"testing"
)

// TestDeepSeekAISummarizer 测试 DeepSeek API 集成
func TestDeepSeekAISummarizer(t *testing.T) {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("无法加载配置: %v", err)
	}

	// 跳过如果没有设置 API 密钥
	if config.DeepSeekAPIKey == "" {
		t.Skip("跳过测试：未设置 DeepSeek API 密钥")
	}

	// 创建 DeepSeek 摘要器
	summarizer := NewDeepSeekAISummarizer(config.DeepSeekAPIKey)

	// 测试内容
	content := `我们介绍了一种名为 AutoGen 的框架，用于使用多个自治 AI 代理构建 AI 应用程序，这些代理可以通过对话进行互动。AutoGen 框架定义了对话式代理的接口，并实现了一组这样的代理，这些代理可以在人类参与的情况下或自主模式下进行交互。AutoGen 的对话式代理可以直接使用，也可以通过修改组合配置轻松自定义，以适应不同的应用需求。`

	// 调用摘要生成
	summary, err := summarizer.Summarize(content)
	if err != nil {
		t.Fatalf("生成摘要失败: %v", err)
	}

	// 验证结果
	if summary == "" {
		t.Error("生成的摘要为空")
	}

	t.Logf("生成的摘要: %s", summary)
}

// TestOpenAIAISummarizer 测试 OpenAI API 集成
func TestOpenAIAISummarizer(t *testing.T) {
	// 加载配置
	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("无法加载配置: %v", err)
	}

	// 跳过如果没有设置 API 密钥
	if config.OpenAIAPIKey == "" {
		t.Skip("跳过测试：未设置 OpenAI API 密钥")
	}

	// 创建 OpenAI 摘要器
	summarizer := NewOpenAISummarizer(config.OpenAIAPIKey)

	// 测试内容
	content := `我们介绍了一种名为 AutoGen 的框架，用于使用多个自治 AI 代理构建 AI 应用程序，这些代理可以通过对话进行互动。AutoGen 框架定义了对话式代理的接口，并实现了一组这样的代理，这些代理可以在人类参与的情况下或自主模式下进行交互。AutoGen 的对话式代理可以直接使用，也可以通过修改组合配置轻松自定义，以适应不同的应用需求。`

	// 调用摘要生成
	summary, err := summarizer.Summarize(content)
	if err != nil {
		t.Fatalf("生成摘要失败: %v", err)
	}

	// 验证结果
	if summary == "" {
		t.Error("生成的摘要为空")
	}

	t.Logf("生成的摘要: %s", summary)
}

// TestMockDeepSeekAISummarizer 测试模拟 DeepSeek 实现
func TestMockDeepSeekAISummarizer(t *testing.T) {
	// 创建配置
	config := &Config{
		AIProvider:     "deepseek",
		DeepSeekAPIKey: "mock_api_key",
	}

	// 创建摘要器
	summarizer, err := NewAISummarizer(config)
	if err != nil {
		t.Fatalf("创建摘要器失败: %v", err)
	}

	// 模拟实际调用
	t.Logf("注意：这个测试不会真正调用 DeepSeek API")

	// 由于真实的 API 调用会失败（mock_api_key 不是有效的密钥），
	// 所以我们不执行 Summarize，而只是验证初始化是否成功
	_, ok := summarizer.(*DeepSeekAISummarizer)
	if !ok {
		t.Errorf("创建的摘要器类型不正确，应该是 *DeepSeekAISummarizer")
	}
}

// TestMockAISummarizer 测试模拟 AI 摘要器
func TestMockAISummarizer(t *testing.T) {
	// 创建配置
	config := &Config{
		AIProvider: "mock",
	}

	// 创建摘要器
	summarizer, err := NewAISummarizer(config)
	if err != nil {
		t.Fatalf("创建摘要器失败: %v", err)
	}

	// 调用摘要生成
	summary, err := summarizer.Summarize("任意内容")
	if err != nil {
		t.Fatalf("生成摘要失败: %v", err)
	}

	// 验证结果
	if summary == "" {
		t.Error("生成的摘要为空")
	}

	t.Logf("生成的摘要: %s", summary)
}
