# Insight-V: Exploring Long-Chain Visual Reasoning with Multimodal Large Language Models

**URL**: http://arxiv.org/abs/2411.14432v2

## 原始摘要

Large Language Models (LLMs) demonstrate enhanced capabilities and
reliability by reasoning more, evolving from Chain-of-Thought prompting to
product-level solutions like OpenAI o1. Despite various efforts to improve LLM
reasoning, high-quality long-chain reasoning data and optimized training
pipelines still remain inadequately explored in vision-language tasks. In this
paper, we present Insight-V, an early effort to 1) scalably produce long and
robust reasoning data for complex multi-modal tasks, and 2) an effective
training pipeline to enhance the reasoning capabilities of multi-modal large
language models (MLLMs). Specifically, to create long and structured reasoning
data without human labor, we design a two-step pipeline with a progressive
strategy to generate sufficiently long and diverse reasoning paths and a
multi-granularity assessment method to ensure data quality. We observe that
directly supervising MLLMs with such long and complex reasoning data will not
yield ideal reasoning ability. To tackle this problem, we design a multi-agent
system consisting of a reasoning agent dedicated to performing long-chain
reasoning and a summary agent trained to judge and summarize reasoning results.
We further incorporate an iterative DPO algorithm to enhance the reasoning
agent's generation stability and quality. Based on the popular LLaVA-NeXT model
and our stronger base MLLM, we demonstrate significant performance gains across
challenging multi-modal benchmarks requiring visual reasoning. Benefiting from
our multi-agent system, Insight-V can also easily maintain or improve
performance on perception-focused multi-modal tasks.


## AI 摘要

本文提出Insight-V框架，旨在提升多模态大语言模型(MLLM)的推理能力。该框架包含两个创新：1) 通过渐进式策略自动生成高质量的长链视觉语言推理数据；2) 设计由推理代理和总结代理组成的多智能体系统，结合迭代DPO算法优化训练流程。实验表明，基于LLaVA-NeXT改进的Insight-V在需要视觉推理的任务上表现显著提升，同时保持感知任务的性能。该方法无需人工标注即可生成结构化长推理数据，并通过多粒度评估确保数据质量，为复杂多模态任务提供了有效的训练解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T04:03:21Z
- **目录日期**: 2025-05-05
