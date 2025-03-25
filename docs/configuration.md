# 配置指南

AI Workflow 提供灵活的配置选项，支持多种配置方式，方便您根据自己的需求进行定制。本文档详细介绍了所有可用的配置选项及其用法。

## 配置方式

本项目支持三种配置方式，按优先级从高到低排列：

1. **命令行参数** - 执行时直接指定的参数，优先级最高
2. **环境变量** - 系统中设置的环境变量
3. **配置文件** - 项目根目录中的 `.env` 文件

## 基本配置

### 数据源配置

| 配置项 | 命令行参数 | 环境变量 | 说明 | 默认值 |
|-------|-----------|---------|------|-------|
| RSS URLs | `-rss` | `APP_RSS_URLS` | RSS源URL，多个URL用逗号分隔 | - |
| 最大结果数 | `-max` | `APP_RSS_MAX_RESULTS` | 每个源处理的最大结果数 | `10` |
| 内容源 | `-sources` | `APP_SOURCES` | 内容源类型，如 arxiv,weibo,twitter | `arxiv` |
| 查询关键词 | `-query` | `APP_QUERY` | 搜索查询关键词 | `ai agent` |

### AI 提供者配置

| 配置项 | 命令行参数 | 环境变量 | 说明 | 默认值 |
|-------|-----------|---------|------|-------|
| AI 提供者 | `-ai` | `APP_AI_PROVIDER` | AI模型提供者 (openai, azure, deepseek, mock) | `openai` |
| OpenAI API密钥 | - | `APP_AI_OPENAI_API_KEY` 或 `OPENAI_API_KEY` | OpenAI API密钥 | - |
| OpenAI模型 | - | `APP_AI_OPENAI_MODEL` | 使用的OpenAI模型 | `gpt-4` |
| Azure API密钥 | - | `APP_AI_AZURE_API_KEY` | Azure API密钥 | - |
| Azure端点 | - | `APP_AI_AZURE_ENDPOINT` | Azure端点URL | - |
| DeepSeek API密钥 | - | `APP_AI_DEEPSEEK_API_KEY` | DeepSeek API密钥 | - |
| 自定义提示词 | `-prompt` | `APP_AI_PROMPT_TEMPLATE` | 自定义提示词模板 | 预设模板 |

### 存储配置

| 配置项 | 命令行参数 | 环境变量 | 说明 | 默认值 |
|-------|-----------|---------|------|-------|
| 存储提供者 | `-storage` | `APP_STORAGE_PROVIDER` | 存储后端类型 (markdown, notion, dynamodb) | `markdown` |
| 输出目录 | `-out` | `APP_STORAGE_OUTPUT_DIR` 或 `OUTPUT_DIR` | Markdown输出目录 | `papers` |
| Notion API密钥 | - | `APP_STORAGE_NOTION_API_KEY` 或 `NOTION_API_KEY` | Notion API密钥 | - |
| Notion数据库ID | - | `APP_STORAGE_NOTION_DATABASE_ID` 或 `NOTION_DATABASE_ID` | Notion数据库ID | - |
| 生成索引 | `-generate-index` | `APP_STORAGE_GENERATE_INDEX` | 是否生成索引文件 | `true` |

### AWS配置 (适用于DynamoDB存储)

| 配置项 | 环境变量 | 说明 | 默认值 |
|-------|---------|------|-------|
| AWS区域 | `APP_AWS_REGION` | AWS区域 | `us-east-1` |
| AWS配置文件 | `APP_AWS_PROFILE` | AWS配置文件名称 | `default` |
| DynamoDB表名 | `APP_AWS_DYNAMODB_TABLE` | DynamoDB表名 | `arxiv_papers` |

### 运行模式配置

| 配置项 | 命令行参数 | 环境变量 | 说明 | 默认值 |
|-------|-----------|---------|------|-------|
| 持续运行 | `-continuous` | `APP_RUN_CONTINUOUS` | 持续运行模式 | `false` |
| 运行间隔 | `-interval` | `APP_RUN_INTERVAL_MINUTES` | 持续运行模式下的间隔（分钟） | `60` |
| 仅生成索引 | `-generate-index` | - | 仅生成索引，不处理内容 | `false` |

## 去重系统配置

| 配置项 | 环境变量 | 说明 | 默认值 |
|-------|---------|------|-------|
| 启用去重 | `APP_DEDUPLICATION_ENABLED` | 是否启用去重功能 | `true` |
| 去重策略 | `APP_DEDUPLICATION_STRATEGIES` | 去重策略，支持 url,content_hash,title,similarity | `url,content_hash` |
| 相似度阈值 | `APP_DEDUPLICATION_SIMILARITY_THRESHOLD` | 内容相似度阈值 (0.0-1.0) | `0.85` |
| 缓存大小 | `APP_DEDUPLICATION_CACHE_SIZE` | 去重缓存大小 | `1000` |
| 数据库类型 | `APP_DEDUPLICATION_DATABASE_TYPE` | 去重数据库类型 (sqlite, mysql, dynamodb) | `sqlite` |
| 数据库路径 | `APP_DEDUPLICATION_DATABASE_PATH` | 数据库路径或连接字符串 | `deduplication.db` |

## Hugo站点配置

Hugo站点的配置主要通过 `hugo_site/config.toml` 文件进行，主要配置项包括：

- `baseURL` - 站点基础URL
- `title` - 站点标题
- `theme` - 使用的主题
- 多语言设置
- 主题相关参数

详细的Hugo站点配置请参考Hugo官方文档或项目中的示例配置文件。

## 配置示例

### 基本命令行示例

```bash
# 基本用法
./ai_workflow -ai openai -storage markdown

# 指定特定的RSS源和最大结果数
./ai_workflow -rss https://rsshub.app/arxiv/search_query=abs:ai+agent -max 5

# 持续运行模式，每小时处理一次
./ai_workflow -continuous -interval 60

# 仅生成索引
./ai_workflow -generate-index
```

### 环境变量配置示例

```bash
# 设置基本环境变量
export APP_AI_PROVIDER=openai
export APP_AI_OPENAI_API_KEY=your_api_key_here
export APP_STORAGE_PROVIDER=notion
export APP_STORAGE_NOTION_API_KEY=your_notion_api_key
export APP_STORAGE_NOTION_DATABASE_ID=your_database_id

# 运行应用
./ai_workflow
```

### 配置文件示例 (.env)

```
# RSS 配置
APP_RSS_URLS=https://rsshub.app/arxiv/search_query=abs:ai+agent&start=0&max_results=10,https://rsshub.app/weibo/user/6105753431
APP_RSS_MAX_RESULTS=10

# AI 提供者配置
APP_AI_PROVIDER=openai
APP_AI_OPENAI_API_KEY=your_openai_api_key
APP_AI_OPENAI_MODEL=gpt-4

# 存储提供者配置
APP_STORAGE_PROVIDER=markdown
APP_STORAGE_OUTPUT_DIR=papers
APP_STORAGE_GENERATE_INDEX=true

# 运行模式
APP_RUN_CONTINUOUS=false
APP_RUN_INTERVAL_MINUTES=60

# 去重配置
APP_DEDUPLICATION_ENABLED=true
APP_DEDUPLICATION_STRATEGIES=url,content_hash
APP_DEDUPLICATION_SIMILARITY_THRESHOLD=0.85
APP_DEDUPLICATION_DATABASE_TYPE=sqlite
APP_DEDUPLICATION_DATABASE_PATH=deduplication.db
```

## 高级配置

### 自定义提示词模板

您可以通过 `APP_AI_PROMPT_TEMPLATE` 环境变量或配置文件自定义AI摘要生成的提示词模板。模板中可以使用以下变量：

- `{{.Content}}` - 原始内容
- `{{.Title}}` - 内容标题
- `{{.URL}}` - 内容URL
- `{{.MaxTokens}}` - 最大令牌数

示例：
```
APP_AI_PROMPT_TEMPLATE="请用中文总结以下AI论文，重点关注创新点和实际应用：\n\n标题：{{.Title}}\n\n内容：{{.Content}}\n\n请提供300字以内的简洁摘要。"
```

### 多数据源配置

当配置多个数据源时，可以为每个源指定特定的处理参数：

```
# 主要配置
APP_SOURCES=arxiv,weibo,twitter

# ArXiv特定配置
APP_SOURCE_ARXIV_QUERY=ai agent
APP_SOURCE_ARXIV_MAX_RESULTS=10

# 微博特定配置
APP_SOURCE_WEIBO_USERS=6105753431,1234567890
APP_SOURCE_WEIBO_MAX_RESULTS=5

# Twitter特定配置
APP_SOURCE_TWITTER_SEARCH=ai research
APP_SOURCE_TWITTER_MAX_RESULTS=20
```

## 故障排除

如果遇到配置问题，请尝试以下方法：

1. 使用 `-verbose` 命令行参数启用详细日志
2. 检查环境变量是否正确设置
3. 确认API密钥和访问令牌是否有效
4. 查看详细日志输出以识别具体问题

如需更多帮助，请在GitHub Issues中提问或参考项目Wiki。 