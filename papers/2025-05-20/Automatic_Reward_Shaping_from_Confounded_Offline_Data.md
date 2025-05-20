# Automatic Reward Shaping from Confounded Offline Data

**URL**: http://arxiv.org/abs/2505.11478v1

## 原始摘要

A key task in Artificial Intelligence is learning effective policies for
controlling agents in unknown environments to optimize performance measures.
Off-policy learning methods, like Q-learning, allow learners to make optimal
decisions based on past experiences. This paper studies off-policy learning
from biased data in complex and high-dimensional domains where \emph{unobserved
confounding} cannot be ruled out a priori. Building on the well-celebrated Deep
Q-Network (DQN), we propose a novel deep reinforcement learning algorithm
robust to confounding biases in observed data. Specifically, our algorithm
attempts to find a safe policy for the worst-case environment compatible with
the observations. We apply our method to twelve confounded Atari games, and
find that it consistently dominates the standard DQN in all games where the
observed input to the behavioral and target policies mismatch and unobserved
confounders exist.


## AI 摘要

该论文针对强化学习中的离策略学习问题，提出了一种能处理观测数据中存在未观测混杂偏差的深度Q学习算法。基于经典深度Q网络(DQN)，新算法通过寻找与观测数据兼容的最坏情况下环境的安全策略，增强了对混杂偏差的鲁棒性。在12个存在混杂因素的Atari游戏测试中，该算法在行为策略与目标策略输入不匹配且存在未观测混杂因素的所有游戏中，均显著优于标准DQN。研究为解决高维复杂领域中的混杂偏差问题提供了有效方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T01:29:05Z
- **目录日期**: 2025-05-20
