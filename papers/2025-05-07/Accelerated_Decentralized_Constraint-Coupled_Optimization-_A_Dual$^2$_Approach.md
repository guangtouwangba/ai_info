# Accelerated Decentralized Constraint-Coupled Optimization: A Dual$^2$ Approach

**URL**: http://arxiv.org/abs/2505.03719v1

## 原始摘要

In this paper, we focus on a class of decentralized constraint-coupled
optimization problem: $\min_{x_i \in \mathbb{R}^{d_i}, i \in \mathcal{I}; y \in
\mathbb{R}^p}$ $\sum_{i=1}^n\left(f_i(x_i) + g_i(x_i)\right) + h(y) \
\text{s.t.} \ \sum_{i=1}^{n}A_ix_i = y$, over an undirected and connected
network of $n$ agents. Here, $f_i$, $g_i$, and $A_i$ represent private
information of agent $i \in \mathcal{I} = \{1, \cdots, n\}$, while $h$ is
public for all agents. Building on a novel dual$^2$ approach, we develop two
accelerated algorithms for solving this problem: the inexact Dual$^2$
Accelerated (iD2A) gradient method and the Multi-consensus inexact Dual$^2$
Accelerated (MiD2A) gradient method. We demonstrate that both iD2A and MiD2A
can guarantee asymptotic convergence under a milder condition on $h$ compared
to existing algorithms. Furthermore, linear convergence is established under
additional assumptions. By employing specialized saddle-point subproblem
solvers, iD2A and MiD2A attain significantly lower communication and
computational complexities than existing algorithms across various scenarios.
Finally, we conduct several numerical experiments to validate our theoretical
results and to showcase the superior performance of iD2A and MiD2A in practice.


## AI 摘要

本文研究了一类去中心化的约束耦合优化问题，提出两种基于新型双重方法（dual²）的加速算法：iD2A和MiD2A。这些算法在较弱的条件下保证渐近收敛，并在额外假设下实现线性收敛。通过专用鞍点子问题求解器，iD2A和MiD2A在通信和计算复杂度上显著优于现有算法。数值实验验证了理论结果，并展示了其优越性能。该问题涉及n个智能体组成的无向连通网络，其中fi、gi和Ai为各智能体私有信息，h为公共信息。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T23:01:35Z
- **目录日期**: 2025-05-07
