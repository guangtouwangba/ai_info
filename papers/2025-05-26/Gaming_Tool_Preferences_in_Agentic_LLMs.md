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

大型语言模型(LLMs)通过模型上下文协议(MCP)可以访问各种外部工具，但其工具选择机制存在脆弱性。研究发现，对工具描述的细微编辑能显著影响LLMs的选择偏好，某些编辑甚至能使工具使用率提升10倍以上(GPT-4.1和Qwen2-7B)。实验表明，不同模型对描述编辑的响应存在差异。这种现象虽然为开发者提供了推广工具的手段，但也暴露出当前代理型LLMs在工具选择机制上的可靠性问题，亟需建立更稳健的基础架构来确保工具使用的客观性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T05:01:41Z
- **目录日期**: 2025-05-26
