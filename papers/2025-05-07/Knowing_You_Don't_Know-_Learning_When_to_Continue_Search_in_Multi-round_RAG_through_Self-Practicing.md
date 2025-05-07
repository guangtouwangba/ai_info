# Knowing You Don't Know: Learning When to Continue Search in Multi-round RAG through Self-Practicing

**URL**: http://arxiv.org/abs/2505.02811v1

## 原始摘要

Retrieval Augmented Generation (RAG) has shown strong capability in enhancing
language models' knowledge and reducing AI generative hallucinations, driving
its widespread use. However, complex tasks requiring multi-round retrieval
remain challenging, and early attempts tend to be overly optimistic without a
good sense of self-skepticism. Current multi-round RAG systems may continue
searching even when enough information has already been retrieved, or they may
provide incorrect answers without having sufficient information or knowledge.
Existing solutions either require large amounts of expensive human-labeled
process supervision data or lead to subpar performance.
  This paper aims to address these limitations by introducing a new framework,
\textbf{SIM-RAG}, to explicitly enhance RAG systems' self-awareness and
multi-round retrieval capabilities. To train SIM-RAG, we first let a RAG system
self-practice multi-round retrieval, augmenting existing question-answer pairs
with intermediate inner monologue reasoning steps to generate synthetic
training data. For each pair, the system may explore multiple retrieval paths,
which are labeled as successful if they reach the correct answer and
unsuccessful otherwise. Using this data, we train a lightweight information
sufficiency Critic. At inference time, the Critic evaluates whether the RAG
system has retrieved sufficient information at each round, guiding retrieval
decisions and improving system-level self-awareness through in-context
reinforcement learning.
  Experiments across multiple prominent RAG benchmarks show that SIM-RAG is an
effective multi-round RAG solution. Furthermore, this framework is
system-efficient, adding a lightweight component to RAG without requiring
modifications to existing LLMs or search engines, and data-efficient,
eliminating the need for costly human-annotated mid-step retrieval process
supervision data.


## AI 摘要

这篇论文提出了SIM-RAG框架，旨在提升检索增强生成(RAG)系统的自我意识和多轮检索能力。通过让RAG系统自我练习多轮检索并生成合成训练数据，研究者训练了一个轻量级的"信息充分性评判器"。该评判器能在推理时评估每轮检索是否充分，从而优化检索决策。实验表明，SIM-RAG在多个基准测试中表现优异，且具有系统高效性（仅添加轻量组件）和数据高效性（无需昂贵的人工标注数据）。该框架有效解决了当前多轮RAG系统检索过度或不足的问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T01:29:05Z
- **目录日期**: 2025-05-07
