# Gaming Tool Preferences in Agentic LLMs

**URL**: http://arxiv.org/abs/2505.18135v1

## 原始摘要

Large language models (LLMs) can now access a wide range of external tools,
thanks to the Model Context Protocol (MCP). This greatly expands their
abilities as various agents. However, LLMs rely entirely on the text
descriptions of tools to decide which ones to use--a process that is
surprisingly fragile. In this work, we expose a vulnerability in prevalent
tool/function-calling protocols by investigating a series of edits to tool
descriptions, some of which can drastically increase a tool's usage from LLMs
when competing with alternatives. Through controlled experiments, we show that
tools with properly edited descriptions receive over 10 times more usage from
GPT-4.1 and Qwen2.5-7B than tools with original descriptions. We further
evaluate how various edits to tool descriptions perform when competing directly
with one another and how these trends generalize or differ across a broader set
of 10 different models. These phenomenons, while giving developers a powerful
way to promote their tools, underscore the need for a more reliable foundation
for agentic LLMs to select and utilize tools and resources.


## AI 摘要

当前大语言模型(LLMs)通过模型上下文协议(MCP)能访问多种外部工具，但其工具选择完全依赖文本描述，存在严重脆弱性。研究发现，对工具描述的特定编辑能显著影响LLMs的选择偏好：经优化的描述可使GPT-4.1和Qwen2-7B等模型对工具的使用率提升10倍以上。实验测试了10种不同模型，发现这一现象具有普遍性。虽然这为开发者推广工具提供了有效手段，但也暴露出LLMs代理系统在工具选择机制上的重大安全隐患，亟需建立更可靠的工具选择基础架构。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T04:04:33Z
- **目录日期**: 2025-05-26
