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

本文提出Insight-V框架，旨在提升多模态大语言模型(MLLM)的推理能力。通过两步自动化流程生成高质量长链推理数据：渐进式生成多样化推理路径+多粒度质量评估。研究发现直接训练难以有效吸收复杂数据，因此设计了双代理系统：推理代理执行长链推理，总结代理评估结果，并采用迭代DPO算法提升生成稳定性。基于LLaVA-NeXT模型改进后，在视觉推理任务中表现显著提升，同时保持感知任务的性能。该方案首次系统探索了多模态场景下的长链推理数据生成与训练优化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T19:02:12Z
- **目录日期**: 2025-05-05
