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

研究发现，大型语言模型(LLMs)通过工具描述文本选择外部工具时存在严重漏洞。实验表明，经过特定编辑的工具描述可使GPT-4.1和Qwen2-7B等模型调用该工具的概率提升10倍以上。研究人员测试了10种不同模型，发现工具描述修改会显著影响模型选择，这虽然为开发者推广工具提供了途径，但也暴露出当前代理型LLMs在工具选择机制上的脆弱性，亟需建立更可靠的基础架构来确保工具选择的客观性。该研究揭示了现有工具调用协议中的安全隐患。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T02:31:21Z
- **目录日期**: 2025-05-26
