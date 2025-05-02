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

本文提出GIT-STORM模型，通过将传统MLP先验替换为掩码生成先验（如MaskGIT），改进了基于随机Transformer的世界模型（STORM）。该模型在Atari 100k基准测试中显著提升了强化学习性能，并首次将Transformer世界模型应用于连续动作环境。通过状态混合函数整合潜在状态表征和动作，模型成功处理连续控制任务，在DeepMind Control Suite上验证了有效性。研究表明，掩码生成先验能提高世界模型的准确性和强化学习策略的有效性，为研究到实际应用的过渡提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T02:30:03Z
- **目录日期**: 2025-05-02
