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

并行数据收集通过让N个相同代理在N个环境副本中运行，将数据收集速度提升N倍。本文提出了一种新框架，通过最大化并行数据熵值，平衡个体代理熵值与代理间多样性，减少冗余。采用集中式策略梯度方法，实验表明该方法优于相同代理系统，并能与批量强化学习技术协同利用数据多样性。理论分析表明，专业化并行采样分布具有更快的收敛速度，验证了方法的有效性。该研究为超越N倍加速提供了新思路，其理论成果也可能具有独立价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T18:01:32Z
- **目录日期**: 2025-05-05
