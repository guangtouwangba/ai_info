# Cosmos-Reason1: From Physical Common Sense To Embodied Reasoning

**URL**: http://arxiv.org/abs/2503.15558v3

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
large language models, Cosmos-Reason1-7B and Cosmos-Reason1-56B. We curate data
and train our models in two stages: Physical AI supervised fine-tuning (SFT)
and Physical AI reinforcement learning (RL). To evaluate our models, we build
comprehensive benchmarks for physical common sense and embodied reasoning
according to our ontologies. Evaluation results show that Physical AI SFT and
RL bring significant improvements. To facilitate the development of Physical
AI, we make our code and pre-trained models available under the NVIDIA Open
Model License at https://github.com/nvidia-cosmos/cosmos-reason1.


## AI 摘要

本文介绍了Cosmos-Reason1模型，这是一个能够理解物理世界并通过长链思维推理生成适当具身决策（如下一步动作）的物理AI系统。研究定义了物理常识和具身推理等关键能力，采用分层本体表示物理常识（空间、时间、物理知识），二维本体实现跨具身泛化。团队开发了7B和56B参数的多模态大语言模型，通过物理AI监督微调(SFT)和强化学习(RL)两阶段训练。基于本体的评估基准显示该方法显著提升了性能。相关代码和预训练模型已在NVIDIA开源许可下发布。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T02:31:36Z
- **目录日期**: 2025-05-21
