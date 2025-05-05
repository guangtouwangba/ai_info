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

本文介绍了Insight-V，这是一个针对多模态大语言模型（MLLM）的视觉语言推理增强框架。它通过两步流程自动生成长链推理数据：首先生成多样化推理路径，再通过多粒度评估确保质量。为解决长数据直接训练效果不佳的问题，Insight-V采用双代理系统（推理代理执行长链推理，总结代理评估结果）并结合迭代DPO算法提升稳定性。基于LLaVA-NeXT模型改进后，该框架在视觉推理任务中表现显著提升，同时保持感知任务的性能。研究填补了高质量长链推理数据与训练优化的空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T10:02:41Z
- **目录日期**: 2025-05-05
