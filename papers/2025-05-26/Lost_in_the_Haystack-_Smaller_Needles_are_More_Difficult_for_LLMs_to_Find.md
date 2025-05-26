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

研究发现，大语言模型(LLM)在处理"大海捞针"任务时，当相关上下文(金标内容)较短时性能显著下降。实验表明，较短的金标内容会降低模型表现并加剧位置敏感性，这对需要整合分散细粒度信息的智能系统构成挑战。这一现象在三个不同领域(常识、生物医学和数学推理)及七种最先进LLM中普遍存在，与模型规模和架构无关。研究揭示了金标内容长度对长上下文问答任务的关键影响，为设计稳健的上下文感知LLM系统提供了重要见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T06:01:00Z
- **目录日期**: 2025-05-26
