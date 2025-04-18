# 贡献指南

非常感谢您对 AI Workflow 项目的关注！我们欢迎各种形式的贡献，无论是代码、文档还是创意。本指南将帮助您了解如何参与项目开发。

## 🌟 贡献方式

您可以通过以下方式参与项目：

1. **提交问题和建议** - 通过 GitHub Issues 报告 bug 或提出新功能建议
2. **改进文档** - 修复文档错误、添加示例或教程
3. **贡献代码** - 修复 bug、添加新功能或优化现有代码
4. **分享使用案例** - 分享您如何使用该项目的案例和经验
5. **代码审查** - 审查和评论其他贡献者的 Pull Requests

## 📝 提交问题

提交问题时，请尽量提供以下信息：

- 清晰的问题描述
- 重现问题的步骤
- 预期结果和实际结果
- 系统环境和软件版本信息
- 相关的日志或截图

## 🔄 开发工作流

我们采用标准的 GitHub Flow 工作流：

1. Fork 仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 💻 开发环境设置

1. 克隆仓库：
```bash
git clone https://github.com/yourusername/ai_workflow.git
cd ai_workflow
```

2. 安装依赖：
```bash
go mod download
```

3. 创建并配置环境变量文件：
```bash
cp .env.example .env
# 编辑 .env 文件设置必要的配置
```

4. 运行测试确保一切正常：
```bash
go test ./...
```

## 📋 代码规范

请遵循以下代码规范：

- 遵循 Go 编码规范和最佳实践
- 所有代码必须包含适当的注释
- 为新功能编写测试
- 遵循现有的代码风格和架构

## 📚 文档贡献

如果您想改进文档：

1. 更新 README.md 或其他相关文档
2. 创建或改进具体功能的使用示例
3. 为复杂功能添加图表或流程图
4. 纠正错误或更新过时的信息

## 🧪 测试指南

添加新功能或修复 bug 时，请确保：

1. 为新功能编写单元测试
2. 确保所有现有测试通过
3. 对于大型功能，添加集成测试
4. 测试不同的配置和边缘情况

## 📝 提交 Pull Request

提交 Pull Request 时，请：

1. 描述您的更改解决了什么问题或添加了什么功能
2. 链接相关的 issues
3. 确保 CI 测试通过
4. 附上必要的截图或演示

## 🔍 代码审查过程

所有 Pull Requests 将经过以下审查过程：

1. 自动化测试和代码质量检查
2. 项目维护者的代码审查
3. 根据反馈进行必要的更改
4. 最终批准并合并

## 📅 发布周期

本项目遵循以下发布周期：

- 小版本（bug 修复）：每 1-2 周
- 中等版本（新功能）：每 1-2 个月
- 主要版本：根据功能开发路线图决定

## 👨‍💻 社区交流

加入我们的社区：

- GitHub Discussions：讨论想法和问题
- 贡献者会议：每月举行，讨论项目方向和优先事项

## 🙏 行为准则

我们期望所有贡献者遵循友好、尊重和包容的行为准则。任何形式的骚扰或不尊重行为都不会被容忍。详情请参阅我们的[行为准则](CODE_OF_CONDUCT.md)。

---

再次感谢您的贡献！如有任何问题，请随时在 Issues 中提问或联系项目维护者。 