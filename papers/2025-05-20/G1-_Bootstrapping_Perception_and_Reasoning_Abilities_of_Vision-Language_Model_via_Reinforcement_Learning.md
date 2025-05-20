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

视觉语言模型（VLMs）在多模态任务中表现优异，但在交互式视觉环境（如游戏）中的决策能力较弱。为解决这一“知行差距”，研究者提出VLM-Gym——一个强化学习（RL）环境，包含多样化视觉游戏，支持可扩展的多游戏并行训练。通过纯RL驱动的自我进化训练G0模型后，进一步开发了G1模型，结合感知增强的冷启动和RL微调。G1模型在所有游戏中超越其教师模型，并优于Claude-3.7等领先专有模型。研究发现，感知与推理能力在RL训练中相互促进。相关代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T23:01:02Z
- **目录日期**: 2025-05-20
