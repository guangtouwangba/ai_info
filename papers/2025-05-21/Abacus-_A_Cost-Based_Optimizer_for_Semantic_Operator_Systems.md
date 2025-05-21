# Abacus: A Cost-Based Optimizer for Semantic Operator Systems

**URL**: http://arxiv.org/abs/2505.14661v1

## 原始摘要

LLMs enable an exciting new class of data processing applications over large
collections of unstructured documents. Several new programming frameworks have
enabled developers to build these applications by composing them out of
semantic operators: a declarative set of AI-powered data transformations with
natural language specifications. These include LLM-powered maps, filters,
joins, etc. used for document processing tasks such as information extraction,
summarization, and more. While systems of semantic operators have achieved
strong performance on benchmarks, they can be difficult to optimize. An
optimizer for this setting must determine how to physically implement each
semantic operator in a way that optimizes the system globally. Existing
optimizers are limited in the number of optimizations they can apply, and most
(if not all) cannot optimize system quality, cost, or latency subject to
constraint(s) on the other dimensions. In this paper we present Abacus, an
extensible, cost-based optimizer which searches for the best implementation of
a semantic operator system given a (possibly constrained) optimization
objective. Abacus estimates operator performance by leveraging a minimal set of
validation examples and, if available, prior beliefs about operator
performance. We evaluate Abacus on document processing workloads in the
biomedical and legal domains (BioDEX; CUAD) and multi-modal question answering
(MMQA). We demonstrate that systems optimized by Abacus achieve 18.7%-39.2%
better quality and up to 23.6x lower cost and 4.2x lower latency than the next
best system.


## AI 摘要

大语言模型(LLMs)为处理海量非结构化文档提供了新的数据处理应用。新兴编程框架允许开发者通过语义操作符(如LLM驱动的映射、过滤、连接等)构建文档处理应用。然而，现有优化器难以全局优化系统质量、成本和延迟。研究者提出Abacus，一种基于成本的扩展优化器，利用少量验证示例和先验知识评估操作符性能，在生物医学(BioDEX)、法律(CUAD)领域和多模态问答(MMQA)任务中，使系统质量提升18.7%-39.2%，成本降低23.6倍，延迟减少4.2倍。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T06:01:49Z
- **目录日期**: 2025-05-21
