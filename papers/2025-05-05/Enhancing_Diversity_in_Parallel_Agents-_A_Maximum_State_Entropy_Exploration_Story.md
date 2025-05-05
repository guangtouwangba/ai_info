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

并行数据收集通过部署多个相同智能体在环境模拟器中加速数据采集，但本文提出通过策略差异化进一步提升效率。研究者设计了一个最大化数据熵的新框架，平衡个体熵与智能体间多样性以减少冗余。该方法采用集中式策略梯度，实验显示其优于同质智能体系统，并能与批量强化学习技术协同利用数据多样性。理论分析表明，差异化并行采样分布能加快收敛速度，验证了该方法的有效性。这一成果为大规模强化学习应用提供了新的优化思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T02:30:50Z
- **目录日期**: 2025-05-05
