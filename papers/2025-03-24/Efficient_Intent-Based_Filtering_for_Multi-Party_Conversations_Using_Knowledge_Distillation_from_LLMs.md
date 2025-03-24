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

大型语言模型（LLMs）在对话AI中展现出卓越能力，但资源消耗巨大。为此，我们提出了一种经济高效的解决方案，通过知识蒸馏开发意图过滤器，筛选出需要LLM处理的对话片段，而非处理所有内容。该方法结合多种策略创建多样化的多参与者对话数据集，并用目标意图进行标注，进而微调MobileBERT模型用于多标签意图分类。该模型在效率和性能间取得平衡，有效过滤对话片段，仅将相关片段传递给LLM处理，显著降低了操作成本。实验证明了该方法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T03:14:53Z
- **目录日期**: 2025-03-24
