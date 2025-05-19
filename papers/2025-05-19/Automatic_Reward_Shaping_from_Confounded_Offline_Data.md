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

这篇论文针对人工智能中在存在未观测混杂因素的高维复杂环境下，基于有偏数据进行离策略学习的问题，提出了一种新型深度强化学习算法。该算法基于深度Q网络(DQN)，旨在寻找与观测数据兼容的最坏情况下仍安全有效的策略。研究在12个存在混杂因素的Atari游戏上进行测试，结果表明在行为策略与目标策略输入不匹配且存在未观测混杂因素的情况下，新算法始终优于标准DQN。这一方法为解决实际应用中观测数据存在偏差时的策略优化问题提供了有效方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T03:22:58Z
- **目录日期**: 2025-05-19
