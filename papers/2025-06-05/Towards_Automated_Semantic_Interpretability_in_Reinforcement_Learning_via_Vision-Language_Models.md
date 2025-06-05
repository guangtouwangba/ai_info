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

iTRACE是一种自动化的强化学习框架，通过结合预训练视觉语言模型（VLM）和可解释的树模型，实现语义可解释性。它利用VLM提取语义特征，并通过轻量级模型替代VLM以解决计算效率问题，再使用树模型进行策略优化。该方法无需人工标注，克服了传统可解释模型依赖人工定义特征的局限性，同时解决了VLM缺乏动作空间基础和无法直接优化策略的问题。实验表明，iTRACE在性能上匹配基于CNN的策略，并优于使用相同可解释特征的多层感知机基线，生成可验证、语义可解释且符合人类预期的行为。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T00:02:30Z
- **目录日期**: 2025-06-05
