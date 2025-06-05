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

iTRACE是一种自动化的可解释强化学习框架，通过结合视觉语言模型(VLM)和基于树的模型，实现语义可解释性。它利用预训练VLM提取语义特征，并通过可解释树模型映射到策略，避免了传统方法依赖人工标注的局限。为解决VLM在RL循环中计算开销大的问题，iTRACE将其输出蒸馏为轻量级模型。实验表明，iTRACE在使用相同可解释特征时优于MLP基线，性能匹配CNN策略，同时产生可验证、语义可解释且符合人类认知的行为，无需人工标注。该框架解决了VLM缺乏动作空间关联和无法直接优化策略的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T01:29:52Z
- **目录日期**: 2025-06-05
