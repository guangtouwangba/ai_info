# Non-Adversarial Inverse Reinforcement Learning via Successor Feature Matching

**URL**: http://arxiv.org/abs/2411.07007v2

## 原始摘要

In inverse reinforcement learning (IRL), an agent seeks to replicate expert
demonstrations through interactions with the environment. Traditionally, IRL is
treated as an adversarial game, where an adversary searches over reward models,
and a learner optimizes the reward through repeated RL procedures. This
game-solving approach is both computationally expensive and difficult to
stabilize. In this work, we propose a novel approach to IRL by direct policy
optimization: exploiting a linear factorization of the return as the inner
product of successor features and a reward vector, we design an IRL algorithm
by policy gradient descent on the gap between the learner and expert features.
Our non-adversarial method does not require learning a reward function and can
be solved seamlessly with existing actor-critic RL algorithms. Remarkably, our
approach works in state-only settings without expert action labels, a setting
which behavior cloning (BC) cannot solve. Empirical results demonstrate that
our method learns from as few as a single expert demonstration and achieves
improved performance on various control tasks.


## AI 摘要

本文提出了一种新的逆向强化学习（IRL）方法，通过直接优化策略而非传统对抗式训练来模仿专家行为。该方法利用回报的线性分解（后继特征与奖励向量的内积），设计了一种基于策略梯度下降的IRL算法，直接最小化学习器与专家特征间的差距。这种非对抗式方法无需学习奖励函数，可无缝结合现有actor-critic强化学习算法，且在仅含状态信息（无专家动作标签）的场景中依然有效，这是行为克隆（BC）无法实现的。实验表明，该方法仅需单次专家演示即可学习，并在多个控制任务中表现优异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T11:02:50Z
- **目录日期**: 2025-04-23
