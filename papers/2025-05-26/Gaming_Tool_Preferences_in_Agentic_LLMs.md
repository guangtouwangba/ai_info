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

大型语言模型(LLMs)通过模型上下文协议(MCP)可以访问各种外部工具，但其工具选择完全依赖文本描述，存在严重脆弱性。研究发现，对工具描述进行特定编辑能显著提升其被选概率：经过优化的描述在GPT-4.1和Qwen2.5-7B中的使用率可达原始描述的10倍以上。实验评估了不同编辑策略在10个模型中的表现差异，揭示了开发者可通过描述优化来推广工具。这种现象凸显了当前基于描述的代理选择机制不可靠，亟需建立更稳健的工具选择基础架构。(98字)

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T08:02:55Z
- **目录日期**: 2025-05-26
