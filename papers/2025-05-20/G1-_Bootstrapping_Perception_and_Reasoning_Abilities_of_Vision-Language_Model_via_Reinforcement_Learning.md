# G1: Bootstrapping Perception and Reasoning Abilities of Vision-Language Model via Reinforcement Learning

**URL**: http://arxiv.org/abs/2505.13426v1

## 原始摘要

Vision-Language Models (VLMs) excel in many direct multimodal tasks but
struggle to translate this prowess into effective decision-making within
interactive, visually rich environments like games. This ``knowing-doing'' gap
significantly limits their potential as autonomous agents, as leading VLMs
often performing badly in simple games. To address this, we introduce VLM-Gym,
a curated reinforcement learning (RL) environment featuring diverse visual
games with unified interfaces and adjustable, compositional difficulty,
specifically designed for scalable multi-game parallel training. Leveraging
VLM-Gym, we train G0 models using pure RL-driven self-evolution, which
demonstrate emergent perception and reasoning patterns. To further mitigate
challenges arising from game diversity, we develop G1 models. G1 incorporates a
perception-enhanced cold start prior to RL fine-tuning. Our resulting G1 models
consistently surpass their teacher across all games and outperform leading
proprietary models like Claude-3.7-Sonnet-Thinking. Systematic analysis reveals
an intriguing finding: perception and reasoning abilities mutually bootstrap
each other throughout the RL training process. Source code including VLM-Gym
and RL training are released at https://github.com/chenllliang/G1 to foster
future research in advancing VLMs as capable interactive agents.


## AI 摘要

本文介绍了VLM-Gym，一个专为视觉语言模型（VLMs）设计的强化学习环境，用于提升其在视觉丰富环境（如游戏）中的决策能力。通过该环境，研究者训练了G0和G1模型，其中G1模型结合了感知增强的冷启动和强化学习微调，显著提升了性能。实验表明，G1模型在所有游戏中均超越其教师模型，并优于领先的专有模型如Claude-3.7-Sonnet-Thinking。研究还发现，感知和推理能力在强化学习过程中相互促进。相关代码已开源以促进未来研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T16:01:07Z
- **目录日期**: 2025-05-20
