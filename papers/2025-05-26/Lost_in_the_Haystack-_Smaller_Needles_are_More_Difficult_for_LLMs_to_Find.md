# Lost in the Haystack: Smaller Needles are More Difficult for LLMs to Find

**URL**: http://arxiv.org/abs/2505.18148v1

## 原始摘要

Large language models (LLMs) face significant challenges with
needle-in-a-haystack tasks, where relevant information ("the needle") must be
drawn from a large pool of irrelevant context ("the haystack"). Previous
studies have highlighted positional bias and distractor quantity as critical
factors affecting model performance, yet the influence of gold context size has
received little attention. We address this gap by systematically studying how
variations in gold context length impact LLM performance on long-context
question answering tasks. Our experiments reveal that LLM performance drops
sharply when the gold context is shorter, i.e., smaller gold contexts
consistently degrade model performance and amplify positional sensitivity,
posing a major challenge for agentic systems that must integrate scattered,
fine-grained information of varying lengths. This pattern holds across three
diverse domains (general knowledge, biomedical reasoning, and mathematical
reasoning) and seven state-of-the-art LLMs of various sizes and architectures.
Our work provides clear insights to guide the design of robust, context-aware
LLM-driven systems.


## AI 摘要

研究表明，大型语言模型（LLMs）在"大海捞针"任务中表现不佳，尤其是当关键信息（"针"）较短且分散在大量无关内容（"干草堆"）中时。实验发现，关键信息越短，模型性能下降越显著，且位置敏感性增强。这一现象在多个领域（常识、生物医学和数学推理）和七种不同架构的先进LLMs中均存在。研究揭示了关键信息长度对模型表现的重要影响，为设计更稳健的上下文感知LLM系统提供了重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T11:00:58Z
- **目录日期**: 2025-05-26
