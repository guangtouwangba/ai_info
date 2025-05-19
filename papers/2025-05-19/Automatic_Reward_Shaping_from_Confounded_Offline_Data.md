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

该论文提出了一种新型深度强化学习算法，针对存在未观测混杂因素的高维复杂环境中的离策略学习问题。基于深度Q网络(DQN)，该算法通过寻找与观测数据兼容的最坏情况下环境的安全策略，有效解决了观测数据中的混杂偏差问题。在12个存在混杂因素的Atari游戏测试中，该算法在行为策略与目标策略输入不匹配且存在未观测混杂因素的所有游戏中，均显著优于标准DQN方法。研究为存在观测偏差的强化学习提供了鲁棒解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T13:10:49Z
- **目录日期**: 2025-05-19
