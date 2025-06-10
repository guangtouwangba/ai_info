# Automated Capability Discovery via Foundation Model Self-Exploration

**URL**: http://arxiv.org/abs/2502.07577v3

## 原始摘要

Foundation models have become general-purpose assistants, exhibiting diverse
capabilities across numerous domains through training on web-scale data. It
remains challenging to precisely characterize even a fraction of the full
spectrum of these abilities and potential risks in any new model. Existing
evaluation approaches often require significant human effort, and it is taking
increasing effort to design ever harder challenges for more capable models. We
introduce Automated Capability Discovery (ACD), a framework that designates one
foundation model as a scientist to systematically propose open-ended tasks
probing the abilities of a subject model (potentially itself). By combining
frontier models with ideas from the field of open-endedness, ACD automatically
and systematically uncovers a diverse spectrum of surprising capabilities and
failures in the subject model. We demonstrate ACD across a range of foundation
models (including the GPT, Claude, and Llama series), showing that it
automatically generates thousands of distinct tasks, which are then clustered
to reveal dozens of broader capability areas and failure modes, that would be
challenging for any single team to uncover. We further validate our method's
automated scoring with extensive human surveys, observing high agreement
between model-generated and human evaluations. By leveraging foundation models'
ability to both create tasks and self-evaluate, ACD is a significant step
toward scalable, automated evaluation of novel AI systems. All code and
evaluation logs are open-sourced at https://github.com/conglu1997/ACD.


## AI 摘要

本文提出了自动化能力发现（ACD）框架，利用前沿大模型作为"科学家"来自动生成开放式任务，系统性地探测目标模型（包括自身）的能力边界和潜在风险。该方法结合开放式演化思想，在GPT、Claude、Llama等系列模型上验证，可自动生成数千个任务并聚类分析，揭示传统人工评估难以发现的能力维度和失败模式。实验表明模型生成评分与人工评估高度一致。ACD通过大模型的任务创造和自评估能力，为新型AI系统的可扩展自动化评估提供了重要突破。相关代码和评估日志已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T22:02:23Z
- **目录日期**: 2025-06-10
