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

本文研究了离散时间随机线性动态网络系统的有限时域最优协方差控制问题。与现有仅考虑动态解耦智能体的研究不同，我们关注具有动态耦合的互联子系统网络，并考虑局部通信约束。作者提出了一种分布式算法，基于系统级综合(SLS)框架将问题转化为带局部约束的凸优化问题，并通过分解和转换处理不可分情况，最终采用改进的共识ADMM方法实现分布式计算。算法仅需邻域子系统状态信息，并在36节点电力系统中验证了有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T13:12:29Z
- **目录日期**: 2025-04-29
