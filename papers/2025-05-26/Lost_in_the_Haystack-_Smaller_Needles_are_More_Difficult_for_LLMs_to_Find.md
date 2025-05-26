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

研究表明，大语言模型（LLMs）在处理"大海捞针"任务时面临显著挑战，尤其当关键信息（"针"）的上下文较短时。实验发现，关键上下文越短，模型性能下降越明显，且位置敏感性增强。这一现象在多个领域（常识、生物医学和数学推理）和七种不同架构的先进LLM中均存在。研究揭示了黄金上下文长度对模型表现的关键影响，为设计更鲁棒、上下文感知的LLM系统提供了重要指导，尤其对需要整合分散细粒度信息的智能代理系统具有特殊意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T18:00:54Z
- **目录日期**: 2025-05-26
