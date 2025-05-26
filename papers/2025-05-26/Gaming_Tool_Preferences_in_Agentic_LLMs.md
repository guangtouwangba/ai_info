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

研究发现，大型语言模型(LLMs)通过模型上下文协议(MCP)调用外部工具时，完全依赖工具文本描述进行选择，这一过程存在严重漏洞。实验表明，对工具描述进行特定编辑后，GPT-4.1和Qwen2-7B等模型使用该工具的频率可提升10倍以上。研究测试了10种不同模型，发现工具描述编辑效果存在差异。这种现象虽然为开发者推广工具提供了有效手段，但也凸显出现有代理型LLMs在工具选择机制上的脆弱性，亟需建立更可靠的工具选择基础架构。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T06:01:33Z
- **目录日期**: 2025-05-26
