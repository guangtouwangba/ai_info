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

大型语言模型(LLMs)通过模型上下文协议(MCP)可以访问各种外部工具，但其工具选择机制存在漏洞。研究发现，仅修改工具描述就能显著影响LLM的选择倾向，优化后的描述可使GPT-4.1和Qwen2-7B等模型对工具的使用率提升10倍以上。实验测试了10种不同模型，发现工具描述编辑会直接影响使用偏好。这种现象虽为开发者提供了推广工具的手段，但也暴露了当前LLM代理工具选择机制的脆弱性，亟需建立更可靠的工具选择基础架构。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T21:01:17Z
- **目录日期**: 2025-05-26
