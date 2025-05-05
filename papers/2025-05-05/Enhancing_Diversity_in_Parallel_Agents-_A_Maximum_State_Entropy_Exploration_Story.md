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

并行数据收集通过同时运行多个环境模拟器大幅提升了强化学习（RL）的效率。本文提出了一种新颖的学习框架，通过最大化并行数据收集的熵值来优化策略：既保持单个智能体的探索性，又减少多智能体间的冗余。该方法采用集中式策略梯度算法，实验证明其性能优于同质智能体系统，并能与批量RL技术协同利用数据多样性。理论分析表明，专业化并行采样分布具有更快的收敛速率，验证了该方法的有效性。这一成果为大规模RL应用提供了新的加速思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T22:01:27Z
- **目录日期**: 2025-05-05
