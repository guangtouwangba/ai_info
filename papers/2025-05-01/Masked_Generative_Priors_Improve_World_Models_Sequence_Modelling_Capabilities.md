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

深度强化学习（RL）已成为复杂环境中构建智能体的主流方法。基于模型的方法通过预测环境动态的世界模型提升数据效率，其中掩码生成建模（如MaskGIT）展现出更强的序列建模能力。研究团队在STORM架构基础上引入MaskGIT先验，提出GIT-STORM模型，在Atari 100k基准测试中显著提升RL任务性能。同时首次将基于Transformer的世界模型拓展至连续动作环境，通过状态混合函数整合潜在状态与动作，在DeepMind Control Suite中验证了其有效性。该研究为构建更精确的世界模型和高效RL策略开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T23:01:44Z
- **目录日期**: 2025-05-01
