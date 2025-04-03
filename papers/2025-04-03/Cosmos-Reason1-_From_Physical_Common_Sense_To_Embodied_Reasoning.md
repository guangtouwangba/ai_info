# Cosmos-Reason1: From Physical Common Sense To Embodied Reasoning

**URL**: http://arxiv.org/abs/2503.15558v2

## 原始摘要

Physical AI systems need to perceive, understand, and perform complex actions
in the physical world. In this paper, we present the Cosmos-Reason1 models that
can understand the physical world and generate appropriate embodied decisions
(e.g., next step action) in natural language through long chain-of-thought
reasoning processes. We begin by defining key capabilities for Physical AI
reasoning, with a focus on physical common sense and embodied reasoning. To
represent physical common sense, we use a hierarchical ontology that captures
fundamental knowledge about space, time, and physics. For embodied reasoning,
we rely on a two-dimensional ontology that generalizes across different
physical embodiments. Building on these capabilities, we develop two multimodal
large language models, Cosmos-Reason1-8B and Cosmos-Reason1-56B. We curate data
and train our models in four stages: vision pre-training, general supervised
fine-tuning (SFT), Physical AI SFT, and Physical AI reinforcement learning (RL)
as the post-training. To evaluate our models, we build comprehensive benchmarks
for physical common sense and embodied reasoning according to our ontologies.
Evaluation results show that Physical AI SFT and reinforcement learning bring
significant improvements. To facilitate the development of Physical AI, we will
make our code and pre-trained models available under the NVIDIA Open Model
License at https://github.com/nvidia-cosmos/cosmos-reason1.


## AI 摘要

该研究提出了Cosmos-Reason1模型，通过分层本体表示物理常识（空间、时间、物理法则）和二维本体实现具身推理，开发了8B和56B参数的多模态大语言模型。模型训练分为视觉预训练、通用微调、物理AI微调和强化学习四个阶段。评估显示物理AI微调和强化学习显著提升了性能。研究建立了基于本体的物理常识和具身推理评测基准，相关代码和预训练模型将通过NVIDIA开放许可发布。该工作为物理AI系统在现实世界中感知、理解和执行复杂动作提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T18:02:08Z
- **目录日期**: 2025-04-03
