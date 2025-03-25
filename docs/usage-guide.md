# 使用指南

本指南将帮助您了解如何充分利用 AI Workflow 平台的各项功能，以满足不同的使用场景需求。

## 目录

- [基本使用](#基本使用)
- [常见使用场景](#常见使用场景)
- [高级应用](#高级应用)
- [最佳实践](#最佳实践)
- [常见问题](#常见问题)

## 基本使用

### 安装和设置

如果您尚未设置 AI Workflow，请按照以下步骤进行：

1. 克隆仓库并进入项目目录：
```bash
git clone https://github.com/yourusername/ai_workflow.git
cd ai_workflow
```

2. 安装依赖：
```bash
go mod download
```

3. 复制示例配置文件：
```bash
cp .env.example .env
```

4. 编辑 `.env` 文件，设置必要的配置项，尤其是 API 密钥。

### 基本命令

最简单的运行方式是：

```bash
go run . -ai openai -storage markdown
```

或者编译后运行：

```bash
go build -o ai_workflow
./ai_workflow -ai openai -storage markdown
```

这将使用 OpenAI 处理内容并将结果保存为 Markdown 文件。

### 查看结果

处理完成后，结果将保存在 `papers/` 目录（默认设置）。您可以：

1. 浏览 `papers/README.md` 获取索引
2. 查看具体日期目录下的各个文件
3. 启动 Hugo 站点预览结果：
```bash
cd hugo_site
hugo server -D
```

然后在浏览器中访问 `http://localhost:1313` 查看内容。

## 常见使用场景

### 场景一：学术研究跟踪

适合研究人员跟踪特定领域的最新论文。

**配置示例**：
```bash
./ai_workflow -ai openai -storage markdown -query "reinforcement learning agent" -max 20
```

**定制建议**：
- 调整 `APP_AI_PROMPT_TEMPLATE` 以突出研究方法和结果
- 使用 `APP_DEDUPLICATION_STRATEGIES=title,similarity` 避免相似研究的重复
- 考虑添加多个特定的 RSS 源来覆盖不同的学术平台

### 场景二：企业情报监控

适合企业监控竞争对手和行业动态。

**配置示例**：
```bash
./ai_workflow -sources arxiv,twitter,weibo -ai deepseek -storage notion -continuous
```

**定制建议**：
- 配置特定的社交媒体账号和关键词
- 将内容保存到 Notion 数据库以便团队协作
- 设置自动化通知，当发现重要情报时通知团队

### 场景三：个人知识库构建

适合个人收集和整理 AI 领域的知识。

**配置示例**：
```bash
./ai_workflow -ai openai -storage markdown -max 5 -interval 1440
```

**定制建议**：
- 使用自定义提示词，针对个人兴趣提取关键信息
- 配置 Hugo 站点，添加个性化分类和标签
- 考虑将站点部署到 GitHub Pages 或 Vercel，方便从任何设备访问

### 场景四：内容创作辅助

适合内容创作者获取写作灵感和素材。

**配置示例**：
```bash
./ai_workflow -sources twitter,weibo,tech_blogs -ai gpt4 -prompt custom_for_content
```

**定制建议**：
- 创建专注于提取叙事要素和趋势的提示词
- 添加特定行业或主题的 RSS 源
- 配置高相似度阈值以获取多样化的内容

### 场景五：团队协作与分享

适合团队共享和讨论最新研究。

**配置示例**：
```bash
./ai_workflow -ai openai -storage notion,markdown -sources arxiv,blogs
```

**定制建议**：
- 同时使用 Notion（团队协作）和 Markdown（存档）
- 配置 GitHub Actions 自动更新共享知识库
- 设置特定的标签体系，便于团队成员筛选感兴趣的内容

## 高级应用

### 自定义数据源

除了预设的数据源外，您还可以添加自定义数据源：

1. 实现 `ContentSource` 接口
2. 在 `sources.go` 中注册您的源
3. 通过配置文件或环境变量设置源特定的参数

**示例**：添加特定的技术博客作为源
```go
type TechBlogSource struct {
    URLs []string
}

func (s *TechBlogSource) Fetch() ([]Content, error) {
    // 实现爬取逻辑
}
```

### 内容处理管道

您可以创建自定义的内容处理管道，在 AI 摘要前后添加处理步骤：

1. 前处理：清理 HTML、提取关键部分、翻译等
2. 后处理：格式化、添加元数据、分类等

通过修改 `processor.go` 文件实现自定义处理步骤。

### 部署到服务器

对于持续运行的场景，建议将应用部署到服务器：

1. 设置系统服务或使用 Docker
2. 配置日志轮转和监控
3. 设置自动备份机制

**Docker示例**：
```dockerfile
FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN go build -o ai_workflow

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/ai_workflow /app/
COPY .env.example /app/.env
RUN apt-get update && apt-get install -y ca-certificates

CMD ["./ai_workflow", "-continuous"]
```

### 与其他工具集成

AI Workflow 可以与多种工具集成：

- **Slack/Discord**：设置通知机制
- **Zapier/IFTTT**：触发其他自动化流程
- **Obsidian/Logseq**：导入到个人知识管理系统
- **自定义API**：开发API层以供其他应用调用

## 最佳实践

### 优化 AI 提示词

提示词设计对于生成高质量摘要至关重要：

- 明确指出您需要的信息类型（方法、结果、应用等）
- 指定输出格式（段落、要点、结构化等）
- 针对特定主题定制提示词
- 考虑添加示例以引导输出风格

**示例提示词**：
```
请用中文总结以下AI论文，采用以下结构：
1. 研究问题和背景（50字）
2. 主要方法和创新点（100字）
3. 实验结果和性能（50字）
4. 实际应用和影响（50字）

标题：{{.Title}}
内容：{{.Content}}

请确保摘要客观、准确，总长度不超过300字。
```

### 有效管理存储

随着时间推移，内容会不断累积：

- 定期备份重要内容
- 考虑使用分层存储策略（活跃内容保存在Notion，历史内容归档到DynamoDB）
- 为较旧的内容设置压缩或归档策略
- 定期维护Hugo站点以保持性能

### 扩展和定制化

根据特定需求扩展系统：

- 添加自定义分析和报告功能
- 开发专门的内容过滤器
- 创建特定领域的处理逻辑
- 自定义Hugo主题以优化阅读体验

### 资源管理

确保系统高效运行：

- 合理设置处理间隔，避免API请求限制
- 监控API使用量和费用
- 对于大规模部署，考虑使用队列系统处理任务
- 定期检查和优化性能瓶颈

## 常见问题

### 问题：摘要质量不理想

**解决方案**：
- 调整提示词，增加具体指导和示例
- 尝试不同的AI模型（GPT-4通常比GPT-3.5效果更好）
- 添加前处理步骤，提取原文中最重要的部分
- 针对特定类型的内容设计专门的提示词

### 问题：系统处理速度慢

**解决方案**：
- 减少每次处理的内容量
- 优化去重策略，减少不必要的处理
- 使用批处理方式调用API
- 考虑并行处理多个数据源

### 问题：内容重复或缺失

**解决方案**：
- 检查去重配置，调整相似度阈值
- 验证RSS源是否正常工作
- 检查API调用是否成功完成
- 查看日志以识别潜在问题

### 问题：部署到GitHub Actions后不工作

**解决方案**：
- 确认所有必要的Secrets已正确设置
- 检查工作流文件格式是否正确
- 审查Action运行日志查找错误
- 验证权限设置是否允许自动提交

---

我们希望本指南能帮助您充分利用AI Workflow的功能。如有其他问题或需要进一步的帮助，请参考项目文档或在GitHub上提交Issue。 