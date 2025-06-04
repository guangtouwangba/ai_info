# Towards Automated Semantic Interpretability in Reinforcement Learning via Vision-Language Models

**URL**: http://arxiv.org/abs/2503.16724v2

## 原始摘要

Semantic interpretability in Reinforcement Learning (RL) enables transparency
and verifiability by making the agent's decisions understandable and
verifiable. Achieving this, however, requires a feature space composed of
human-understandable concepts, which traditionally rely on human specification
and may fail to generalize to unseen environments. We introduce interpretable
Tree-based Reinforcement learning via Automated Concept Extraction (iTRACE), an
automated framework that leverages pre-trained vision-language models (VLM) for
semantic feature extraction and interpretable tree-based models for policy
optimization. iTRACE first extracts semantically meaningful features, then maps
them to policies via interpretable trees. To address the impracticality of
running VLMs in RL loops, we distill their outputs into a lightweight model. By
leveraging Vision-Language Models (VLMs) to automate tree-based reinforcement
learning, iTRACE eliminates the need for human annotation traditionally
required by interpretable models, while also addressing the limitations of VLMs
alone, such as their lack of grounding in action spaces and inability to
directly optimize policies. iTRACE outperforms MLP baselines that use the same
interpretable features and matches the performance of CNN-based policies,
producing verifiable, semantically interpretable, and human-aligned behaviors
without requiring human annotation.


## AI 摘要

iTRACE是一个自动化强化学习框架，通过结合视觉语言模型(VLM)和可解释的树模型，实现语义可解释的决策。它利用预训练VLM提取语义特征，并蒸馏为轻量模型以避免实时计算开销，再通过树模型映射到策略。相比传统依赖人工标注的方法，iTRACE自动生成可理解的概念特征，在保持与CNN策略相当性能的同时，产生可验证、语义可解释且与人类认知一致的行为。该框架解决了VLM缺乏动作空间关联和直接策略优化能力的问题，无需人工标注即可实现泛化，超越了使用相同特征的MLP基线模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T22:02:10Z
- **目录日期**: 2025-06-04
