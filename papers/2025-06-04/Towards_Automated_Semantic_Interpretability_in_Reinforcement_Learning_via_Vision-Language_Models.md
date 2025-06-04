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

iTRACE是一种自动化的强化学习框架，通过结合预训练视觉语言模型(VLM)和可解释的树模型，实现语义可解释的决策。它利用VLM自动提取语义特征，并通过轻量级模型蒸馏VLM输出，避免实时推理的开销。树模型将这些特征映射为可验证的策略，无需人工标注。iTRACE在性能上匹配CNN策略，同时产生语义可解释、人类可验证的行为，解决了传统可解释模型依赖人工标注和VLM缺乏动作空间基础的问题。该框架实现了自动化、高性能且可解释的强化学习，超越了传统MLP基线。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T18:02:22Z
- **目录日期**: 2025-06-04
