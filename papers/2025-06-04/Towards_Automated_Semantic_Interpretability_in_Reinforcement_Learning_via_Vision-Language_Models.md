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

iTRACE是一个自动化强化学习框架，通过结合预训练视觉语言模型(VLM)和可解释树模型，实现语义可解释的决策。它利用VLM自动提取语义特征，并蒸馏为轻量模型以解决计算瓶颈，再通过可解释树模型映射到策略。相比传统方法，iTRACE无需人工标注，同时克服了VLM无法直接优化策略的局限。实验表明，其性能与CNN策略相当，优于使用相同特征的MLP基线，能产生可验证、语义可解释且符合人类预期的行为。该框架为强化学习提供了自动化的可解释性解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T05:02:29Z
- **目录日期**: 2025-06-04
