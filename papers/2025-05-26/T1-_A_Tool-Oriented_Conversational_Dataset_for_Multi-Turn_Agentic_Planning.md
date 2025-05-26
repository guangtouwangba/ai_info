# T1: A Tool-Oriented Conversational Dataset for Multi-Turn Agentic Planning

**URL**: http://arxiv.org/abs/2505.16986v1

## 原始摘要

Large Language Models (LLMs) have demonstrated impressive capabilities as
intelligent agents capable of solving complex problems. However, effective
planning in scenarios involving dependencies between API or tool
calls-particularly in multi-turn conversations-remains a significant challenge.
To address this, we introduce T1, a tool-augmented, multi-domain, multi-turn
conversational dataset specifically designed to capture and manage inter-tool
dependencies across diverse domains. T1 enables rigorous evaluation of agents'
ability to coordinate tool use across nine distinct domains (4 single domain
and 5 multi-domain) with the help of an integrated caching mechanism for both
short- and long-term memory, while supporting dynamic replanning-such as
deciding whether to recompute or reuse cached results. Beyond facilitating
research on tool use and planning, T1 also serves as a benchmark for evaluating
the performance of open-source language models. We present results powered by
T1-Agent, highlighting their ability to plan and reason in complex,
tool-dependent scenarios.


## AI 摘要

本文介绍了T1数据集，这是一个支持多领域、多轮对话的工具增强数据集，旨在解决大语言模型（LLMs）在API或工具调用依赖关系中的规划挑战。T1覆盖9个领域（4个单领域和5个多领域），通过集成缓存机制（短期/长期记忆）和动态重规划能力（如复用缓存或重新计算），评估智能体在多工具协同任务中的表现。该数据集不仅推动工具使用与规划研究，还可作为开源语言模型的基准测试平台。实验通过T1-Agent验证了模型在复杂工具依赖场景下的规划与推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T01:29:35Z
- **目录日期**: 2025-05-26
