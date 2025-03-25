# 🚀 AI Workflow - 智能研究与知识管理平台

<p align="center">
  <img src="https://img.shields.io/github/stars/siqiuchen/ai_workflow?style=social" alt="GitHub Stars">
  <img src="https://img.shields.io/github/forks/siqiuchen/ai_workflow?style=social" alt="GitHub Forks">
  <img src="https://img.shields.io/github/issues/siqiuchen/ai_workflow" alt="GitHub Issues">
  <img src="https://img.shields.io/github/license/siqiuchen/ai_workflow" alt="License">
</p>

<p align="center">
  <b>自动获取、整理并可视化AI领域最新研究与进展的一站式解决方案</b>
</p>

## 🌟 项目概览

AI Workflow 是一个综合性平台，帮助研究人员、开发者和AI爱好者保持对AI领域最新进展的了解。该平台整合了多种功能，包括自动化数据收集、智能摘要生成、可视化展示和知识库管理。

### 主要功能模块：

1. **📊 智能内容采集** - 从多个来源（ArXiv、技术博客、社交媒体等）自动获取最新AI相关内容
2. **🧠 AI驱动摘要** - 使用先进的AI模型自动生成中英文摘要
3. **🌐 知识库构建** - 构建结构化的AI研究知识库，支持多种存储方式
4. **📱 多平台展示** - 通过Hugo静态网站、Markdown文档等多种方式展示内容
5. **🔄 自动化工作流** - 利用GitHub Actions实现全自动化数据收集与更新

## ✨ 为什么选择 AI Workflow？

- **节省时间** - 自动跟踪AI领域的最新发展，无需手动搜索和整理
- **深度见解** - AI生成的摘要提供关键信息，帮助快速理解复杂概念
- **知识沉淀** - 构建个人或团队的AI知识库，促进学习与分享
- **无缝集成** - 支持多种平台和工具，适应不同的工作流程
- **开源灵活** - 完全开源，可根据需求自定义和扩展

## 🖼️ 平台展示

_这里可以添加平台截图或演示GIF_

## 🚀 快速开始

### 前提条件

- Go 1.21 或更高版本
- 可选：OpenAI、Azure或DeepSeek API密钥（用于AI摘要生成）
- 可选：Notion API密钥（如使用Notion存储）

### 安装

```bash
git clone https://github.com/siqiuchen/ai_workflow
cd ai_workflow
go mod download
```

### 配置

本项目支持三种配置方式，优先级从高到低：

1. 命令行参数
2. 环境变量
3. 配置文件 (.env)

#### 使用配置文件

复制 `.env.example` 文件并重命名为 `.env`：

```bash
cp .env.example .env
```

编辑 `.env` 文件，根据需要设置参数（详细配置选项请参考 [配置文档](docs/configuration.md)）。

### 使用方式

#### 1. 内容采集与处理

```bash
# 使用OpenAI处理ArXiv论文
./ai_workflow -ai openai -storage markdown -max 5

# 从多个来源采集内容
./ai_workflow -sources arxiv,weibo,twitter -ai deepseek -storage notion
```

#### 2. 启动Hugo站点预览

```bash
cd hugo_site
hugo server -D
```

## 📚 核心功能详解

### 内容采集系统

AI Workflow支持从多种来源采集内容：

- **学术论文** - 从ArXiv、IEEE等学术平台获取最新研究论文
- **技术博客** - 跟踪主要AI公司和研究机构的技术博客
- **社交媒体** - 监控Twitter、微博等平台上的AI相关讨论
- **新闻资讯** - 收集AI领域的最新新闻报道

### AI摘要引擎

该系统使用先进的大语言模型为采集的内容生成高质量摘要：

- **多模型支持** - 支持OpenAI、Azure、DeepSeek等多种大语言模型
- **双语摘要** - 自动生成中英文摘要，满足不同读者需求
- **关键点提取** - 智能识别并突出内容中的关键点和创新点
- **可定制提示词** - 灵活调整摘要生成的风格和侧重点

### 知识库管理

支持多种存储和管理方式：

- **结构化存储** - 按日期、主题、来源等组织内容
- **多后端支持** - Markdown、Notion、DynamoDB等多种存储方式
- **自动索引** - 自动生成索引文件，便于内容导航
- **标签系统** - 通过标签对内容进行分类和检索

### Hugo静态网站

内置Hugo网站提供美观的Web界面：

- **响应式设计** - 适配桌面和移动设备
- **多语言支持** - 支持中英文界面切换
- **主题定制** - 基于PaperMod主题，可进一步定制界面
- **搜索功能** - 内置全文搜索，快速找到所需内容

### 自动化工作流

通过GitHub Actions实现全自动化运行：

- **定时执行** - 定期自动收集和更新内容
- **自动提交** - 自动将新内容提交到仓库
- **部署集成** - 可配置自动部署到GitHub Pages、Vercel等平台
- **通知机制** - 支持邮件、Discord等渠道的更新通知

## 🛠️ 高级定制

### 添加新的数据源

在 `sources.go` 文件中实现 `ContentSource` 接口：

```go
type NewSource struct {
    // 配置...
}

func (s *NewSource) Fetch() ([]Content, error) {
    // 实现内容获取逻辑...
}
```

### 添加新的AI提供者

在 `ai.go` 文件中实现 `AISummarizer` 接口：

```go
type NewAIProvider struct {
    // 配置...
}

func (p *NewAIProvider) Summarize(content string) (string, error) {
    // 实现摘要生成逻辑...
}
```

### 自定义Hugo主题

1. 修改 `hugo_site/config.toml` 文件中的主题设置
2. 在 `hugo_site/layouts` 目录中创建自定义模板
3. 在 `hugo_site/static/css` 中定制样式

## 📊 数据分析扩展 (实验性)

> 新增功能：利用收集的数据进行高级分析和可视化

- **趋势分析** - 追踪AI研究热点的变化趋势
- **关联发现** - 识别不同研究方向之间的关联
- **作者网络** - 构建研究人员和机构的合作网络
- **引用分析** - 追踪论文引用关系和影响力

## 🤝 参与贡献

我们欢迎各种形式的贡献，包括但不限于：

- 报告问题和提出建议
- 提交代码改进和新功能
- 改进文档和添加教程
- 分享使用经验和案例

请查看 [贡献指南](CONTRIBUTING.md) 了解更多详情。

## 📜 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

## 🙏 鸣谢

- 感谢所有开源社区的贡献者
- 特别感谢为本项目提供灵感和支持的所有人

---

<p align="center">
  <a href="https://github.com/siqiuchen/ai_workflow">GitHub</a> •
  <a href="https://siqiuchen.github.io/ai_workflow">文档</a> •
  <a href="https://github.com/siqiuchen/ai_workflow/issues">问题反馈</a>
</p>

# ArXiv AI Agent Paper Crawler

这是一个自动获取和总结 ArXiv 上 AI Agent 相关论文的工具。该工具会定期从 ArXiv 获取最新的 AI Agent 相关论文，使用 AI 生成中文摘要，并将结果保存到指定位置（如 Notion、Markdown 文件等）。

## 功能特点

- 自动从 ArXiv 获取最新的 AI Agent 相关论文
- 使用 AI（如 OpenAI、Azure、DeepSeek）生成中文摘要
- 支持多种存储后端（Notion、Markdown 文件）
- 通过 GitHub Actions 实现定时自动运行
- 模块化的设计，可轻松扩展 AI 和存储提供者
- 灵活的配置系统，支持环境变量、配置文件和命令行参数
- 内置强大的去重系统，支持多种去重策略和数据库后端

## 快速开始

### 前提条件

- Go 1.21 或更高版本
- 如使用 OpenAI：OpenAI API 密钥
- 如使用 Notion：Notion API 密钥和数据库 ID

### 安装

```bash
git clone https://github.com/user/arxiv-ai-agent
cd arxiv-ai-agent
go mod download
```

### 配置

本项目支持三种配置方式，优先级从高到低：

1. 命令行参数
2. 环境变量
3. 配置文件 (.env)

#### 使用配置文件

复制 `.env.example` 文件并重命名为 `.env`：

```bash
cp .env.example .env
```

然后编辑 `.env` 文件，根据需要设置参数：

```env
# RSS 配置
# 多个URL用逗号分隔
APP_RSS_URLS=https://rsshub.app/arxiv/search_query=abs:ai+agent&start=0&max_results=10&sortBy=lastUpdatedDate,https://rsshub.app/weibo/user/6105753431
APP_RSS_MAX_RESULTS=10

# AI 提供者配置
APP_AI_PROVIDER=openai
APP_AI_OPENAI_API_KEY=your_openai_api_key

# 存储提供者配置
APP_STORAGE_PROVIDER=markdown
APP_STORAGE_OUTPUT_DIR=papers

# 运行模式
APP_RUN_CONTINUOUS=false

# 去重配置
APP_DEDUPLICATION_ENABLED=true
# 可选：url, content_hash, title, similarity
APP_DEDUPLICATION_STRATEGIES=url,content_hash
# 内容相似度阈值 (0.0-1.0)，仅当使用similarity策略时有效
APP_DEDUPLICATION_SIMILARITY_THRESHOLD=0.85
# 去重缓存大小
APP_DEDUPLICATION_CACHE_SIZE=1000
# 数据库类型：sqlite, mysql, dynamodb (可扩展)
APP_DEDUPLICATION_DATABASE_TYPE=sqlite
# 去重数据库路径或连接字符串
APP_DEDUPLICATION_DATABASE_PATH=deduplication.db
```

#### 使用环境变量

项目支持两种环境变量命名格式：

1. 应用特定格式（推荐）：
```bash
export APP_AI_PROVIDER=openai
export APP_AI_OPENAI_API_KEY=your_openai_api_key
export APP_STORAGE_PROVIDER=markdown
```

2. 兼容旧版本：
```bash
export OPENAI_API_KEY=your_openai_api_key
export NOTION_API_KEY=your_notion_api_key
export NOTION_DATABASE_ID=your_notion_database_id
export OUTPUT_DIR=path_to_output_directory
```

#### 使用命令行参数

```bash
go run . -ai openai -storage markdown -max 5 -continuous
```

### 运行

```bash
go run .
```

或者先编译再运行：

```bash
go build -o arxiv-agent
./arxiv-agent
```

### 命令行参数

- `-rss`: 单个RSS源URL（如果提供，将覆盖配置文件中的所有URL）
- `-max`: 每个RSS源处理的最大结果数（默认为10）
- `-ai`: AI 提供者（`openai`、`azure`、`deepseek` 或 `mock`，默认为 `openai`）
- `-storage`: 存储提供者（`notion` 或 `markdown`，默认为 `markdown`）
- `-continuous`: 持续运行模式，每小时执行一次（默认为单次运行）

示例:
```bash
# 使用单一RSS源
./arxiv-agent -rss https://rsshub.app/arxiv/search_query=abs:ai+agent -ai deepseek -storage markdown -max 5 -continuous

# 使用配置文件中的多个RSS源
./arxiv-agent -ai deepseek -storage markdown -max 5
```

## 部署到 GitHub Actions

1. Fork 本仓库
2. 在仓库设置中添加以下 Secrets：
   - `APP_AI_OPENAI_API_KEY`：OpenAI API 密钥
   - `APP_STORAGE_NOTION_API_KEY`：Notion API 密钥（如使用Notion存储）
   - `APP_STORAGE_NOTION_DATABASE_ID`：Notion 数据库 ID（如使用Notion存储）
3. GitHub Actions 会根据 `.github/workflows/arxiv-crawler.yml` 中的设置定时运行

### 自动提交功能

程序通过GitHub Actions运行后，会自动：

1. 收集和处理论文数据
2. 将结果保存到按日期组织的目录中（例如 `papers/2023-06-01/`）
3. 生成或更新索引文件 (`papers/README.md`)
4. 自动提交更改到仓库
5. 在提交消息中包含当前日期（例如 "📚 添加 2023-06-01 新论文和更新索引"）

你不需要手动提交任何内容 - GitHub Actions 会自动处理一切。这使得你可以：

- 每小时自动收集和整理最新的研究论文
- 在GitHub上浏览按日期组织的论文，方便查找最新内容
- 通过索引文件快速访问所有论文

### 自定义配置

你可以在GitHub仓库设置中添加以下密钥来自定义行为：

- `APP_STORAGE_OUTPUT_DIR`: 自定义输出目录（默认为 "papers"）
- `APP_RSS_URLS`: 自定义RSS源（多个URL用逗号分隔）
- `APP_DEDUPLICATION_ENABLED`: 控制是否启用去重功能（true/false）

## 项目结构

- `main.go`: 程序入口
- `config.go`: 配置管理
- `rss.go`: RSS 获取相关
- `ai.go`: AI 服务相关
- `storage.go`: 存储服务相关

## 扩展

### 添加新的 AI 提供者

在 `ai.go` 文件中，通过实现 `AISummarizer` 接口并更新 `NewAISummarizer` 工厂函数来添加新的 AI 提供者：

```go
// 创建新的 AI 提供者
type NewAIProvider struct {
    // 配置...
}

// 实现 Summarize 方法
func (p *NewAIProvider) Summarize(content string) (string, error) {
    // 实现...
}

// 在工厂函数中添加
func NewAISummarizer(provider string) (AISummarizer, error) {
    switch provider {
    // 现有的 case...
    case "new-provider":
        // 返回新的提供者
        return NewNewAIProvider(), nil
    }
}
```

### 添加新的存储提供者

在 `storage.go` 文件中，通过实现 `StorageProvider` 接口并更新 `NewStorageProvider` 工厂函数来添加新的存储提供者：

```go
// 创建新的存储提供者
type NewStorageProvider struct {
    // 配置...
}

// 实现 Save 方法
func (p *NewStorageProvider) Save(paper Paper) error {
    // 实现...
}

// 在工厂函数中添加
func NewStorageProvider(provider string) (StorageProvider, error) {
    switch provider {
    // 现有的 case...
    case "new-storage":
        // 返回新的提供者
        return NewNewStorageProvider(), nil
    }
}
```

## 运行测试

该项目包含完整的测试套件，用于验证工作流和各个组件的功能。

### 快速测试

使用提供的测试脚本运行所有测试并查看结果：

```bash
# 赋予脚本执行权限
chmod +x run_test.sh

# 运行测试脚本
./run_test.sh
```

### 手动运行测试

也可以手动运行特定的测试：

```bash
# 运行所有测试
go test -v ./...

# 运行特定测试
go test -v -run TestWorkflowWithMarkdown

# 运行特定测试文件
go test -v ./workflow_test.go
```

### 测试说明

测试套件包含以下主要测试：

1. **工作流测试** - 测试完整的工作流程，包括 RSS 获取、AI 摘要和文件存储
2. **Markdown 存储测试** - 专门测试 Markdown 文件存储功能
3. **端到端测试** - 测试整个应用程序的端到端流程

测试使用模拟（Mock）组件避免对外部服务的依赖，确保测试可以在没有 API 密钥的情况下运行。

## 许可证

MIT 

# 存储系统

本项目支持多种存储后端，您可以选择最适合您需求的：

## Markdown 存储

最简单的存储方式，将数据保存为本地Markdown文件：

```env
APP_STORAGE_PROVIDER=markdown
APP_STORAGE_OUTPUT_DIR=papers
```

Markdown存储提供者现在支持以下功能：

1. **按日期组织内容** - 所有论文按照保存日期自动组织到日期目录中（格式：YYYY-MM-DD）
2. **自动生成索引** - 在`papers/README.md`中生成索引文件，列出所有可用的日期和论文
3. **GitHub Actions集成** - 与GitHub Actions无缝集成，自动提交新内容

你也可以使用命令行标志来单独生成索引：

```bash
./arxiv-agent -generate-index
```

### 目录结构示例

```
papers/
├── README.md                    # 自动生成的索引文件
├── 2023-06-01/                  # 日期目录
│   ├── Paper_Title_1.md
│   └── Paper_Title_2.md
└── 2023-06-02/                  # 最新的日期目录
    ├── Paper_Title_3.md
    └── Paper_Title_4.md
```

## Notion 存储

将数据保存到Notion数据库，需要配置Notion API：

```env
APP_STORAGE_PROVIDER=notion
APP_STORAGE_NOTION_API_KEY=your_notion_api_key_here
APP_STORAGE_NOTION_DATABASE_ID=your_notion_database_id_here
```

## DynamoDB 存储

将数据保存到AWS DynamoDB，适合大规模部署和高可用性需求：

```env
APP_STORAGE_PROVIDER=dynamodb
APP_AWS_REGION=us-east-1  # 您的AWS区域
APP_AWS_PROFILE=default   # 您的AWS配置文件
APP_AWS_DYNAMODB_TABLE=arxiv_papers  # 表名
```

### DynamoDB 存储设置

1. 安装必要的依赖：

```bash
# 运行安装脚本
chmod +x install_aws.sh
./install_aws.sh

# 或手动安装依赖
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
```

2. 确保您已配置AWS凭证：
   - 通过AWS CLI：`aws configure`
   - 或在`~/.aws/credentials`文件中设置

3. 首次使用时，系统会自动创建必要的DynamoDB表和索引。

4. 表结构设计：
   - 主键：`url`（用于唯一标识每篇论文）
   - 全局二级索引：`saved_at-index`（用于按保存日期查询）

### DynamoDB优势

- **高可用性**：AWS管理的服务，提供高可用性和可靠性
- **可扩展性**：自动扩展，无需担心存储容量
- **云存储**：数据安全存储在云端，可从任何地方访问
- **查询能力**：支持按URL或保存日期查询
- **与AWS生态集成**：可与其他AWS服务无缝集成

# 去重系统

本项目实现了一个强大的去重系统，可以避免重复处理相同的内容。

## 去重特性

- **多策略支持**：同时支持多种去重策略，包括URL去重、内容哈希去重、标题去重和内容相似度去重
- **多数据库支持**：支持多种数据库作为后端存储，目前包括SQLite、MySQL和DynamoDB
- **高效缓存**：使用LRU缓存提高去重检查的性能
- **可配置的相似度阈值**：对于内容相似度去重，可以配置相似度阈值

## 去重配置

在配置文件中添加以下设置：

```env
# 去重配置
APP_DEDUPLICATION_ENABLED=true
# 可选：url, content_hash, title, similarity
APP_DEDUPLICATION_STRATEGIES=url,content_hash
# 内容相似度阈值 (0.0-1.0)，仅当使用similarity策略时有效
APP_DEDUPLICATION_SIMILARITY_THRESHOLD=0.85
# 去重缓存大小
APP_DEDUPLICATION_CACHE_SIZE=1000
# 数据库类型：sqlite, mysql, dynamodb (可扩展)
APP_DEDUPLICATION_DATABASE_TYPE=sqlite
# 去重数据库路径或连接字符串
APP_DEDUPLICATION_DATABASE_PATH=deduplication.db
```

## 数据库配置

### SQLite (默认)

SQLite是默认的数据库后端，不需要额外的配置：

```env
APP_DEDUPLICATION_DATABASE_TYPE=sqlite
APP_DEDUPLICATION_DATABASE_PATH=deduplication.db
```

### MySQL

要使用MySQL作为数据库后端，你需要：

1. 安装MySQL驱动：

```bash
go get github.com/go-sql-driver/mysql
```

2. 在`deduplication.go`文件中取消注释MySQL驱动导入行：

```go
import (
    // ...其他导入
    _ "github.com/mattn/go-sqlite3"
    // 取消下面这行的注释
    _ "github.com/go-sql-driver/mysql"
)
```

3. 在配置文件中设置MySQL连接：

```env
APP_DEDUPLICATION_DATABASE_TYPE=mysql
APP_DEDUPLICATION_DATABASE_PATH=mysql://user:password@localhost:3306/deduplication
```

MySQL连接字符串格式：`mysql://用户名:密码@主机名:端口/数据库名`

### DynamoDB

要使用AWS DynamoDB作为去重系统的数据库后端，需要：

1. 安装AWS SDK依赖（与存储相同的依赖）：

```bash
# 运行安装脚本
chmod +x install_aws.sh
./install_aws.sh
```

2. 在配置文件中设置DynamoDB：

```env
APP_DEDUPLICATION_DATABASE_TYPE=dynamodb
# 使用与存储相同的AWS配置
APP_AWS_REGION=us-east-1
APP_AWS_PROFILE=default
# 可选：设置表前缀
APP_DEDUPLICATION_DATABASE_PATH=dedup_
```

使用DynamoDB作为去重系统的优势：
- **跨实例共享去重信息**：多个应用实例可以共享同一份去重数据
- **持久化存储**：无需担心本地数据库损坏
- **无限扩展**：可以处理大量的去重记录
- **与存储系统统一**：如果已经使用DynamoDB作为存储，可以统一数据库技术栈

**注意**：使用DynamoDB作为去重系统将会创建以下表：
- `{前缀}url_index`：URL去重表
- `{前缀}content_hash_index`：内容哈希去重表
- `{前缀}title_index`：标题去重表
- `{前缀}content_index`：内容相似度去重表

## 添加自定义数据库提供者

该系统设计为可扩展的，你可以通过实现`DBProvider`接口添加对其他数据库的支持：

```go
// DBProvider 定义数据库提供者接口
type DBProvider interface {
    // 初始化数据库，创建必要的表和索引
    Initialize() error
    
    // 关闭数据库连接
    Close() error
    
    // 各种数据操作方法...
}
```

然后在`NewDBProvider`函数中添加对新数据库的支持。