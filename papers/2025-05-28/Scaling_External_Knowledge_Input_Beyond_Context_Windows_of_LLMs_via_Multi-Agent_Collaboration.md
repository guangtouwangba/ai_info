# Scaling External Knowledge Input Beyond Context Windows of LLMs via Multi-Agent Collaboration

**URL**: http://arxiv.org/abs/2505.21471v1

## 原始摘要

With the rapid advancement of post-training techniques for reasoning and
information seeking, large language models (LLMs) can incorporate a large
quantity of retrieved knowledge to solve complex tasks. However, the limited
context window of LLMs obstructs scaling the amount of external knowledge
input, prohibiting further improvement, especially for tasks requiring
significant amount of external knowledge. Existing context window extension
methods inevitably cause information loss. LLM-based multi-agent methods emerge
as a new paradigm to handle massive input in a distributional manner, where we
identify two core bottlenecks in existing knowledge synchronization and
reasoning processes. In this work, we develop a multi-agent framework,
$\textbf{ExtAgents}$, to overcome the bottlenecks and enable better scalability
in inference-time knowledge integration without longer-context training.
Benchmarked with our enhanced multi-hop question answering test,
$\textbf{$\boldsymbol{\infty}$Bench+}$, and other public test sets including
long survey generation, ExtAgents significantly enhances the performance over
existing non-training methods with the same amount of external knowledge input,
regardless of whether it falls $\textit{within or exceeds the context window}$.
Moreover, the method maintains high efficiency due to high parallelism. Further
study in the coordination of LLM agents on increasing external knowledge input
could benefit real-world applications.


## AI 摘要

这篇论文提出了一个名为ExtAgents的多智能体框架，旨在解决大语言模型(LLMs)在知识整合和推理过程中面临的两个核心瓶颈问题。通过分布式处理大量外部知识输入，该框架无需长上下文训练即可实现更好的推理时知识扩展能力。在增强的多跳问答测试集∞Bench+和其他公开测试集上的实验表明，ExtAgents显著优于现有非训练方法，无论输入知识是否超出上下文窗口限制，同时保持高并行性带来的高效性。该研究为实际应用中处理大规模外部知识提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T08:02:35Z
- **目录日期**: 2025-05-28
