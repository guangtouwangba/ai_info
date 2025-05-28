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

这篇论文提出了一个名为ExtAgents的多智能体框架，旨在解决大语言模型(LLMs)在处理大量外部知识时的上下文窗口限制问题。通过创新的知识同步和推理机制，该框架能在不增加上下文训练长度的情况下，有效整合推理时的大量外部知识。在增强的多跳问答测试集∞Bench+和其他公开测试集上的实验表明，ExtAgents显著优于现有非训练方法，无论输入知识是否超出上下文窗口。该方法还保持了高并行性带来的高效性，为需要整合大量外部知识的实际应用提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T10:02:18Z
- **目录日期**: 2025-05-28
