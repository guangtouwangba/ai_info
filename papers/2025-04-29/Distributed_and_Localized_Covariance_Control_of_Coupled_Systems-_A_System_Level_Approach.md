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

本文提出了一种分布式算法，用于解决动态耦合网络系统的有限时域最优协方差控制问题。与现有仅研究动态解耦代理的工作不同，该方法针对受局部通信约束的互联子系统。基于系统级综合框架，首先将局部协方差控制问题转化为带局部约束的凸优化问题；然后利用部分可分离结构将其分解为更小的子问题，并通过变换处理不可分离情况；最后采用改进的共识ADMM方法，根据各子系统的局部信息和通信约束进行分布式计算。在36个互联子系统的电网模型上验证了算法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T08:02:32Z
- **目录日期**: 2025-04-29
