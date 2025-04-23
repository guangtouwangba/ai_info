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

本文提出了一种新的逆向强化学习(IRL)方法，通过直接优化策略来避免传统对抗性方法的计算复杂性和稳定性问题。该方法利用回报的线性分解特性，将问题转化为对学习者与专家特征差距的策略梯度下降优化。这种非对抗性方法无需学习奖励函数，可直接与现有actor-critic强化学习算法结合使用，且仅需专家状态信息而无需动作标签。实验表明，该方法能从单个专家演示中学习，并在多个控制任务上取得更好性能。相比需要动作标签的行为克隆方法，该算法在仅状态信息场景下表现突出。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T21:02:39Z
- **目录日期**: 2025-04-23
