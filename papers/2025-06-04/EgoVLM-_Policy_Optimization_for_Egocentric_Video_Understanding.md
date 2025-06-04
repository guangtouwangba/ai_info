# EgoVLM: Policy Optimization for Egocentric Video Understanding

**URL**: http://arxiv.org/abs/2506.03097v1

## 原始摘要

Emerging embodied AI applications, such as wearable cameras and autonomous
agents, have underscored the need for robust reasoning from first person video
streams. We introduce EgoVLM, a vision-language model specifically designed to
integrate visual comprehension and spatial-temporal reasoning within egocentric
video contexts. EgoVLM is fine-tuned via Group Relative Policy Optimization
(GRPO), a reinforcement learning method adapted to align model outputs with
human-like reasoning steps. Following DeepSeek R1-Zero's approach, we directly
tune using RL without any supervised fine-tuning phase on chain-of-thought
(CoT) data. We evaluate EgoVLM on egocentric video question answering
benchmarks and show that domain-specific training substantially improves
performance over general-purpose VLMs. Our EgoVLM-3B, trained exclusively on
non-CoT egocentric data, outperforms the base Qwen2.5-VL 3B and 7B models by
14.33 and 13.87 accuracy points on the EgoSchema benchmark, respectively. By
explicitly generating reasoning traces, EgoVLM enhances interpretability,
making it well-suited for downstream applications. Furthermore, we introduce a
novel keyframe-based reward that incorporates salient frame selection to guide
reinforcement learning optimization. This reward formulation opens a promising
avenue for future exploration in temporally grounded egocentric reasoning.


## AI 摘要

研究人员开发了EgoVLM，一种专为第一人称视频设计的视觉语言模型，通过强化学习（GRPO方法）优化，无需监督微调即可生成类人推理步骤。在EgoSchema基准测试中，3B参数的EgoVLM（仅用非链式推理数据训练）性能显著超越通用VLMs，比Qwen2.5-VL 3B/7B模型分别高出14.33和13.87个准确点。该模型通过显式生成推理轨迹增强可解释性，并创新性地采用基于关键帧的奖励机制，利用显著帧选择指导强化学习优化，为时序第一人称推理研究开辟了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T23:01:56Z
- **目录日期**: 2025-06-04
