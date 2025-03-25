# Efficient Intent-Based Filtering for Multi-Party Conversations Using Knowledge Distillation from LLMs

**URL**: http://arxiv.org/abs/2503.17336v1

## 原始摘要

Large language models (LLMs) have showcased remarkable capabilities in
conversational AI, enabling open-domain responses in chat-bots, as well as
advanced processing of conversations like summarization, intent classification,
and insights generation. However, these models are resource-intensive,
demanding substantial memory and computational power. To address this, we
propose a cost-effective solution that filters conversational snippets of
interest for LLM processing, tailored to the target downstream application,
rather than processing every snippet. In this work, we introduce an innovative
approach that leverages knowledge distillation from LLMs to develop an
intent-based filter for multi-party conversations, optimized for compute power
constrained environments. Our method combines different strategies to create a
diverse multi-party conversational dataset, that is annotated with the target
intents and is then used to fine-tune the MobileBERT model for multi-label
intent classification. This model achieves a balance between efficiency and
performance, effectively filtering conversation snippets based on their
intents. By passing only the relevant snippets to the LLM for further
processing, our approach significantly reduces overall operational costs
depending on the intents and the data distribution as demonstrated in our
experiments.


## AI 摘要

本文提出了一种低成本解决方案，通过知识蒸馏从大型语言模型(LLM)中提取意图过滤器，用于多轮对话处理。该方法首先构建多样化的多轮对话数据集，标注目标意图后微调MobileBERT模型进行多标签意图分类。该模型在效率和性能间取得平衡，能有效筛选出需要LLM处理的对话片段。实验表明，通过仅将相关片段传递给LLM处理，可显著降低整体运算成本。这种方法特别适合计算资源受限的环境，为LLM的对话处理提供了一种经济高效的替代方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T00:01:36Z
- **目录日期**: 2025-03-25
