# Masked Generative Priors Improve World Models Sequence Modelling Capabilities

**URL**: http://arxiv.org/abs/2410.07836v5

## 原始摘要

Deep Reinforcement Learning (RL) has become the leading approach for creating
artificial agents in complex environments. Model-based approaches, which are RL
methods with world models that predict environment dynamics, are among the most
promising directions for improving data efficiency, forming a critical step
toward bridging the gap between research and real-world deployment. In
particular, world models enhance sample efficiency by learning in imagination,
which involves training a generative sequence model of the environment in a
self-supervised manner. Recently, Masked Generative Modelling has emerged as a
more efficient and superior inductive bias for modelling and generating token
sequences. Building on the Efficient Stochastic Transformer-based World Models
(STORM) architecture, we replace the traditional MLP prior with a Masked
Generative Prior (e.g., MaskGIT Prior) and introduce GIT-STORM. We evaluate our
model on two downstream tasks: reinforcement learning and video prediction.
GIT-STORM demonstrates substantial performance gains in RL tasks on the Atari
100k benchmark. Moreover, we apply Transformer-based World Models to continuous
action environments for the first time, addressing a significant gap in prior
research. To achieve this, we employ a state mixer function that integrates
latent state representations with actions, enabling our model to handle
continuous control tasks. We validate this approach through qualitative and
quantitative analyses on the DeepMind Control Suite, showcasing the
effectiveness of Transformer-based World Models in this new domain. Our results
highlight the versatility and efficacy of the MaskGIT dynamics prior, paving
the way for more accurate world models and effective RL policies.


## AI 摘要

本文提出GIT-STORM模型，将传统MLP先验替换为掩码生成先验（如MaskGIT），显著提升了基于Transformer的世界模型性能。在Atari 100k基准测试中，该模型在强化学习任务上表现优异，并首次成功应用于连续动作环境（如DeepMind Control Suite），通过状态混合函数整合潜在状态与动作。实验表明，掩码生成建模作为更高效的归纳偏置，能有效提升样本效率，推动世界模型向实际应用迈进。研究验证了Transformer世界模型在新领域的适用性，为开发更精确的环境模型和强化学习策略开辟了新途径。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T19:01:46Z
- **目录日期**: 2025-05-01
