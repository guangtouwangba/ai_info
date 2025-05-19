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

这篇论文提出了一种新型深度强化学习算法，用于解决复杂高维环境中存在未观测混杂因素的偏置数据问题。基于深度Q网络(DQN)，该算法通过寻找与观测数据兼容的最坏情况下安全策略，增强了对混杂偏置的鲁棒性。在12个存在混杂因素的Atari游戏测试中，当行为策略和目标策略输入不匹配时，该算法在所有游戏中都显著优于标准DQN。研究结果表明，该方法能有效处理观测数据中的混杂偏置，为存在未观测混杂因素的强化学习任务提供了更可靠的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T16:01:20Z
- **目录日期**: 2025-05-19
