---
title: "PLAY2PROMPT: Zero-shot Tool Instruction Optimization for LLM Agents via Tool Play"
date: 2025-03-20T03:11:00+00:00
source_url: "http://arxiv.org/abs/2503.14432v1"
categories: ["2025-03-20"]
tags: ["Zero-shot", "Tool", "Instruction", "Optimization", "LLM"]
summary: "为了解决大型语言模型（LLMs）在零样本工具使用中的挑战，研究者提出了PLAY2PROMPT框架。该框架通过自动“试错”探索工具的输入输出行为，无需标注数据即可优化工具文档并生成使用示例。这些示例不仅指导LLM推理，还作为验证进一步提升工具使用效果。实验表明，PLAY2PROMPT显著提升了开放和封闭模型在零样本任务中的表现，为领域特定工具集成提供了可扩展且有效的解决方案。"
---

# PLAY2PROMPT: Zero-shot Tool Instruction Optimization for LLM Agents via Tool Play

**原始链接**: [查看原文](http://arxiv.org/abs/2503.14432v1)

## 原始摘要

Large language models (LLMs) are increasingly integrated with specialized
external tools, yet many tasks demand zero-shot tool usage with minimal or
noisy documentation. Existing solutions rely on manual rewriting or labeled
data for validation, making them inapplicable in true zero-shot settings. To
address these challenges, we propose PLAY2PROMPT, an automated framework that
systematically "plays" with each tool to explore its input-output behaviors.
Through this iterative trial-and-error process, PLAY2PROMPT refines tool
documentation and generates usage examples without any labeled data. These
examples not only guide LLM inference but also serve as validation to further
enhance tool utilization. Extensive experiments on real-world tasks demonstrate
that PLAY2PROMPT significantly improves zero-shot tool performance across both
open and closed models, offering a scalable and effective solution for
domain-specific tool integration.

## AI 摘要

为了解决大型语言模型（LLMs）在零样本工具使用中的挑战，研究者提出了PLAY2PROMPT框架。该框架通过自动“试错”探索工具的输入输出行为，无需标注数据即可优化工具文档并生成使用示例。这些示例不仅指导LLM推理，还作为验证进一步提升工具使用效果。实验表明，PLAY2PROMPT显著提升了开放和封闭模型在零样本任务中的表现，为领域特定工具集成提供了可扩展且有效的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T03:11:13Z
- **目录日期**: 2025-03-20
