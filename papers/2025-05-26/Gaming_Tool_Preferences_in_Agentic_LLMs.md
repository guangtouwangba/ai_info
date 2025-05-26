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

研究发现，大型语言模型（LLMs）通过工具描述文本选择外部工具时存在严重脆弱性。实验显示，对工具描述进行特定编辑后，GPT-4.1和Qwen2-7B等模型使用该工具的频率可提升10倍以上。研究测试了10种不同模型，证实工具描述的微小改动会显著影响模型的选择倾向。这种现象虽然为开发者提供了推广工具的手段，但也暴露出当前代理型LLMs在工具选择机制上的安全隐患，亟需建立更可靠的工具选择基础架构。该发现揭示了模型上下文协议（MCP）中亟待解决的关键漏洞。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T20:01:34Z
- **目录日期**: 2025-05-26
