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

研究发现大型语言模型(LLMs)在工具调用上存在严重漏洞：通过修改工具描述文本，可以显著影响模型对工具的选择。实验显示，经过适当编辑的工具描述能让GPT-4.1和Qwen2-7B等模型的使用率提升10倍以上。研究测试了10种不同模型，发现这一现象具有普遍性。虽然这为开发者推广工具提供了便利，但也暴露出当前基于文本描述的工具选择机制存在脆弱性，凸显了需要为智能代理建立更可靠的工具选择基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T01:29:03Z
- **目录日期**: 2025-05-27
