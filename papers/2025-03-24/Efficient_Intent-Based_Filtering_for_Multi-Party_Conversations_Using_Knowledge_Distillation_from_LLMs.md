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

大型语言模型（LLMs）在对话AI中展示了显著能力，如开放域响应、对话摘要、意图分类和洞察生成。然而，这些模型资源消耗大，需要大量内存和计算能力。为此，我们提出了一种成本效益高的解决方案，通过筛选与目标下游应用相关的对话片段进行LLM处理，而非处理所有片段。我们采用知识蒸馏技术，开发了一个针对计算能力受限环境的多方对话意图过滤器，并利用多样化策略创建了标注目标意图的多方对话数据集，用于微调MobileBERT模型进行多标签意图分类。该方法在效率和性能之间取得平衡，显著降低了运营成本。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T05:01:35Z
- **目录日期**: 2025-03-24
