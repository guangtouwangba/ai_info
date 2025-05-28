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

这篇论文提出了一个多智能体框架ExtAgents，旨在解决大语言模型(LLM)在处理海量外部知识时面临的上下文窗口限制问题。通过创新的知识同步和推理机制，该框架无需长上下文训练就能有效整合推理时知识，显著提升了在增强版多跳问答测试∞Bench+等基准上的表现。相比现有非训练方法，ExtAgents在知识量无论是否超出上下文窗口时都表现更优，同时保持了高度并行化的效率优势。研究表明，优化LLM智能体在增加知识输入时的协调机制，将对实际应用产生重要价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T16:02:25Z
- **目录日期**: 2025-05-28
