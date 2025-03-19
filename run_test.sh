#!/bin/bash

# 设置颜色
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}开始测试 ArXiv AI Agent 工作流...${NC}"

# 从.env文件加载环境变量（如果存在）
if [ -f .env ]; then
    echo -e "${GREEN}发现.env文件，加载环境变量...${NC}"
    export $(grep -v '^#' .env | xargs)
fi

# 确保依赖已安装
echo -e "${YELLOW}检查并下载依赖...${NC}"
go mod download

# 运行不依赖外部 API 的测试
echo -e "${YELLOW}运行基本测试...${NC}"
go test -v -run TestMarkdownStorage
go test -v -run TestWorkflowWithMarkdown
go test -v -run TestMainWithMarkdown
go test -v -run TestMockAISummarizer

# 检查是否设置了 API 密钥，运行集成测试
echo -e "${YELLOW}检查是否可以运行集成测试...${NC}"

# 检查新旧两种环境变量名称
if [ -n "$APP_AI_OPENAI_API_KEY" ] || [ -n "$OPENAI_API_KEY" ]; then
    echo -e "${GREEN}找到 OpenAI API 密钥，运行 OpenAI 集成测试...${NC}"
    go test -v -run TestOpenAIAISummarizer
else
    echo -e "${YELLOW}未设置 OpenAI API 密钥，跳过 OpenAI 集成测试${NC}"
fi

if [ -n "$APP_AI_DEEPSEEK_API_KEY" ] || [ -n "$DEEPSEEK_API_KEY" ]; then
    echo -e "${GREEN}找到 DeepSeek API 密钥，运行 DeepSeek 集成测试...${NC}"
    go test -v -run TestDeepSeekAISummarizer
    go test -v -run TestDeepSeekWithMarkdown
else
    echo -e "${YELLOW}未设置 DeepSeek API 密钥，跳过 DeepSeek 集成测试${NC}"
fi

# 打印当前环境变量（仅用于调试）
echo -e "${YELLOW}当前环境变量信息（调试用）:${NC}"
echo "APP_AI_OPENAI_API_KEY存在: $([ -n "$APP_AI_OPENAI_API_KEY" ] && echo '是' || echo '否')"
echo "OPENAI_API_KEY存在: $([ -n "$OPENAI_API_KEY" ] && echo '是' || echo '否')"
echo "APP_AI_DEEPSEEK_API_KEY存在: $([ -n "$APP_AI_DEEPSEEK_API_KEY" ] && echo '是' || echo '否')"
echo "DEEPSEEK_API_KEY存在: $([ -n "$DEEPSEEK_API_KEY" ] && echo '是' || echo '否')"

# 检查是否有测试输出目录
echo -e "${YELLOW}检查生成的 Markdown 文件...${NC}"

if [ -d "test_output" ]; then
    find test_output -name "*.md" -exec echo "文件: {}" \; -exec cat {} \;
fi

if [ -d "test_output_storage" ]; then
    find test_output_storage -name "*.md" -exec echo "文件: {}" \; -exec cat {} \;
fi

if [ -d "test_output_main" ]; then
    find test_output_main -name "*.md" -exec echo "文件: {}" \; -exec cat {} \;
fi

if [ -d "test_output_deepseek" ]; then
    find test_output_deepseek -name "*.md" -exec echo "文件: {}" \; -exec cat {} \;
fi

echo -e "${GREEN}所有测试完成!${NC}"

# 询问是否清理测试生成的文件
echo -e "${YELLOW}是否清理测试生成的文件? (y/n)${NC}"
read -r clean

if [ "$clean" = "y" ] || [ "$clean" = "Y" ]; then
    echo -e "${YELLOW}清理测试文件...${NC}"
    rm -rf test_output test_output_storage test_output_main test_output_deepseek
    echo -e "${GREEN}清理完成!${NC}"
fi 