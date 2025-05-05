# Enhancing Diversity in Parallel Agents: A Maximum State Entropy Exploration Story

**URL**: http://arxiv.org/abs/2505.01336v1

## 原始摘要

Parallel data collection has redefined Reinforcement Learning (RL), unlocking
unprecedented efficiency and powering breakthroughs in large-scale real-world
applications. In this paradigm, $N$ identical agents operate in $N$ replicas of
an environment simulator, accelerating data collection by a factor of $N$. A
critical question arises: \textit{Does specializing the policies of the
parallel agents hold the key to surpass the $N$ factor acceleration?} In this
paper, we introduce a novel learning framework that maximizes the entropy of
collected data in a parallel setting. Our approach carefully balances the
entropy of individual agents with inter-agent diversity, effectively minimizing
redundancies. The latter idea is implemented with a centralized policy gradient
method, which shows promise when evaluated empirically against systems of
identical agents, as well as synergy with batch RL techniques that can exploit
data diversity. Finally, we provide an original concentration analysis that
shows faster rates for specialized parallel sampling distributions, which
supports our methodology and may be of independent interest.


## AI 摘要

并行数据收集通过让N个相同智能体在N个环境副本中运行，将数据收集速度提升N倍。本文提出了一种新型学习框架，通过最大化并行环境中的数据熵来突破N倍加速限制。该方法平衡了单智能体熵与智能体间多样性，减少数据冗余，并采用集中式策略梯度方法。实验表明，该框架优于相同智能体系统，并能与利用数据多样性的批量强化学习技术协同。理论分析表明，专用并行采样分布具有更快的收敛速率，验证了方法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T07:01:33Z
- **目录日期**: 2025-05-05
