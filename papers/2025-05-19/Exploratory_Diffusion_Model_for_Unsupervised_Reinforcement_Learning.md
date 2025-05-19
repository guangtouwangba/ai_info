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

本文提出探索性扩散模型（ExDM），利用扩散模型的强大表达能力拟合无监督强化学习中的探索数据，以促进高效探索并为下游任务提供良好初始化。ExDM通过扩散模型准确估计回放缓冲区中收集数据的分布，并引入基于分数的内在奖励机制，激励智能体探索访问较少的状态。此外，研究提供了扩散策略微调的理论分析和实用算法，解决了多步采样带来的训练不稳定和计算复杂度等关键挑战。实验表明，ExDM在复杂结构环境中显著优于现有方法，实现了更高效的探索和快速的下游任务适应能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T11:02:23Z
- **目录日期**: 2025-05-19
