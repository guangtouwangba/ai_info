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

本文介绍了Insight-V，一个旨在提升多模态大语言模型（MLLM）推理能力的框架。它通过两步流程自动生成长且结构化的多模态推理数据：首先生成多样化的长推理路径，再通过多粒度评估确保数据质量。为解决直接训练效果不佳的问题，Insight-V采用双智能体系统（推理智能体+总结智能体）和迭代DPO算法来优化推理稳定性和质量。基于LLaVA-NeXT及更强基础模型，该方法在视觉推理任务中表现显著提升，同时保持感知任务的性能。研究填补了高质量长链推理数据与训练优化的研究空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T02:31:44Z
- **目录日期**: 2025-05-05
