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

研究发现，大语言模型(LLM)在"大海捞针"任务中表现受关键信息长度影响显著。当相关上下文("金内容")较短时，模型性能急剧下降，且位置敏感性增强。这一现象在三个不同领域(常识、生物医学和数学推理)及七种先进LLM中均成立。研究表明，处理分散的细粒度变长信息是LLM面临的主要挑战，这对需要整合此类信息的智能系统设计具有重要启示。研究填补了此前对关键信息长度影响关注不足的空白，为开发稳健的上下文感知系统提供了指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T03:21:28Z
- **目录日期**: 2025-05-26
