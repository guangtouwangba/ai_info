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

研究发现，大语言模型（LLM）在处理"大海捞针"任务时面临重大挑战，即从大量无关信息中提取少量关键内容。实验表明，当关键信息（"针"）的上下文较短时，LLM性能会显著下降，且位置敏感性增强。这一现象在三个不同领域（常识、生物医学和数学推理）及七种最先进的LLM中都普遍存在。研究指出，关键信息上下文长度是影响模型表现的重要因素，这对需要整合分散、细粒度信息的智能系统设计提出了新挑战。该发现为开发更健壮的上下文感知LLM系统提供了重要指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T17:00:57Z
- **目录日期**: 2025-05-26
