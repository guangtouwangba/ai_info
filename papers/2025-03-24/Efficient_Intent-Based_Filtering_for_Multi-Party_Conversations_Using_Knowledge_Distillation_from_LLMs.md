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

大型语言模型(LLMs)在对话AI中展现出强大能力，但其高资源需求限制了应用。为此，研究者提出了一种经济高效的解决方案：通过知识蒸馏从LLMs中训练出轻量级意图过滤器，仅筛选出与下游任务相关的对话片段供LLM处理。该方法构建了多标签意图分类数据集，并微调MobileBERT模型进行意图识别。实验表明，这种方案能在保证性能的同时显著降低计算成本，特别适合计算资源受限的环境。通过仅处理相关片段，系统整体运行成本得以大幅降低。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T18:01:29Z
- **目录日期**: 2025-03-24
