# Distributed and Localized Covariance Control of Coupled Systems: A System Level Approach

**URL**: http://arxiv.org/abs/2503.02094v2

## 原始摘要

This work is concerned with the finite-horizon optimal covariance steering of
networked systems governed by discrete-time stochastic linear dynamics. In
contrast with existing work that has only considered systems with dynamically
decoupled agents, we consider a dynamically coupled system composed of
interconnected subsystems subject to local communication constraints. In
particular, we propose a distributed algorithm to compute the localized optimal
feedback control policy for each individual subsystem, which depends only on
the local state histories of its neighboring subsystems. Utilizing the
system-level synthesis (SLS) framework, we first recast the localized
covariance steering problem as a convex SLS problem with locality constraints.
Subsequently, exploiting its partially separable structure, we decompose the
latter problem into smaller subproblems, introducing a transformation to deal
with nonseparable instances. Finally, we employ a variation of the consensus
alternating direction method of multipliers (ADMM) to distribute computation
across subsystems on account of their local information and communication
constraints. We demonstrate the effectiveness of our proposed algorithm on a
power system with 36 interconnected subsystems.


## AI 摘要

本文研究了离散时间随机线性动态网络系统的有限时间最优协方差控制问题。与现有仅考虑动态解耦智能体的研究不同，我们关注具有动态耦合特性的互联子系统网络，并考虑局部通信约束。基于系统级综合(SLS)框架，我们将局部化协方差控制问题转化为带局部约束的凸优化问题，并通过分解和变换处理不可分离情况。提出采用改进的共识ADMM算法实现基于局部信息的分布式计算。最后在含36个子系统的电力网络中验证了算法的有效性。该工作为动态耦合网络系统提供了分布式最优控制新方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T20:03:05Z
- **目录日期**: 2025-04-29
