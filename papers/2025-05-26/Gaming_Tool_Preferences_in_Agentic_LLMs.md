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

大型语言模型(LLMs)通过模型上下文协议(MCP)可以访问各种外部工具，但其工具选择完全依赖文本描述，存在明显漏洞。研究发现，适当修改工具描述可使GPT-4.1和Qwen2-7B等模型对某些工具的使用率提升10倍以上。实验评估了不同描述修改方式在10个模型中的表现差异。这种现象虽然为开发者推广工具提供了有效手段，但也凸显出现有代理型LLMs在工具选择和资源利用方面需要更可靠的基础机制。该研究揭示了当前工具调用协议的脆弱性，强调需改进LLMs的工具选择机制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T03:22:01Z
- **目录日期**: 2025-05-26
