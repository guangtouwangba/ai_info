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

研究人员开发了Cosmos-Reason1模型（8B和56B参数），用于物理AI系统的感知、理解和决策。该模型通过长链思维推理生成自然语言决策，基于分层本体表示物理常识（空间、时间、物理定律）和二维本体实现跨物理形态的具身推理。模型训练分四个阶段：视觉预训练、通用微调、物理AI微调和强化学习。评估显示后两阶段显著提升性能。研究团队将按NVIDIA开放许可公开代码和预训练模型（GitHub链接），并建立了基于本体的物理常识与具身推理基准测试体系。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T16:02:07Z
- **目录日期**: 2025-04-03
