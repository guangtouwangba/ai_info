# Exploratory Diffusion Model for Unsupervised Reinforcement Learning

**URL**: http://arxiv.org/abs/2502.07279v2

## 原始摘要

Unsupervised reinforcement learning (URL) aims to pre-train agents by
exploring diverse states or skills in reward-free environments, facilitating
efficient adaptation to downstream tasks. As the agent cannot access extrinsic
rewards during unsupervised exploration, existing methods design intrinsic
rewards to model the explored data and encourage further exploration. However,
the explored data are always heterogeneous, posing the requirements of powerful
representation abilities for both intrinsic reward models and pre-trained
policies. In this work, we propose the Exploratory Diffusion Model (ExDM),
which leverages the strong expressive ability of diffusion models to fit the
explored data, simultaneously boosting exploration and providing an efficient
initialization for downstream tasks. Specifically, ExDM can accurately estimate
the distribution of collected data in the replay buffer with the diffusion
model and introduces the score-based intrinsic reward, encouraging the agent to
explore less-visited states. After obtaining the pre-trained policies, ExDM
enables rapid adaptation to downstream tasks. In detail, we provide theoretical
analyses and practical algorithms for fine-tuning diffusion policies,
addressing key challenges such as training instability and computational
complexity caused by multi-step sampling. Extensive experiments demonstrate
that ExDM outperforms existing SOTA baselines in efficient unsupervised
exploration and fast fine-tuning downstream tasks, especially in structurally
complicated environments.


## AI 摘要

本文提出了一种基于扩散模型的探索性强化学习方法（ExDM），用于无监督环境下的智能体预训练。该方法利用扩散模型强大的表达能力拟合探索数据，同时通过基于分数的内在奖励机制鼓励智能体探索未访问状态。ExDM不仅能准确估计回放缓冲区中数据的分布，还为下游任务提供了高效的初始化策略。研究还针对扩散策略微调中的训练不稳定性和计算复杂度问题提供了理论分析和实用算法。实验表明，ExDM在复杂结构环境中显著优于现有方法，实现了更高效的无监督探索和快速下游任务适应。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T10:02:34Z
- **目录日期**: 2025-05-19
