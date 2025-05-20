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

该研究提出了Cosmos-Reason1模型，旨在通过长链思维推理帮助物理AI系统理解和执行现实世界的复杂任务。研究定义了物理常识和具身推理两大核心能力，分别采用分层本体和二维本体进行知识表示。团队开发了7B和56B两个多模态大语言模型，通过监督微调(SFT)和强化学习(RL)两阶段训练。基于本体的评估基准显示，这两种训练方法显著提升了模型性能。相关代码和预训练模型已在NVIDIA开放模型许可下开源，以推动物理AI发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T17:02:10Z
- **目录日期**: 2025-05-20
