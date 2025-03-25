# 🚀 AI Workflow - 智能研究与知识管理平台

<p align="center">
  <img src="https://img.shields.io/github/stars/guangtouwangba/ai_info?style=social" alt="GitHub Stars">
  <img src="https://img.shields.io/github/forksguangtouwangba/ai_info?style=social" alt="GitHub Forks">
  <img src="https://img.shields.io/github/issues/guangtouwangba/ai_info" alt="GitHub Issues">
  <img src="https://img.shields.io/github/license/guangtouwangba/ai_info" alt="License">
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