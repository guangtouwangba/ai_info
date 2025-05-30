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

大型语言模型（LLMs）在"大海捞针"任务中面临显著挑战，即需要从大量无关信息中提取关键内容。研究发现，当关键上下文（"针"）较短时，模型性能会急剧下降，且位置敏感性增强。这一现象在三个不同领域（常识、生物医学和数学推理）及七种先进LLM中均存在。研究表明，关键上下文的长度变化会显著影响模型表现，这对需要整合分散细粒度信息的智能系统构成主要挑战。该发现为设计鲁棒的上下文感知LLM系统提供了重要指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T01:28:30Z
- **目录日期**: 2025-05-27
