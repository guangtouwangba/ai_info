# Solving Multiagent Path Finding on Highly Centralized Networks

**URL**: http://arxiv.org/abs/2412.09433v2

## 原始摘要

The Mutliagent Path Finding (MAPF) problem consists of identifying the
trajectories that a set of agents should follow inside a given network in order
to reach their desired destinations as soon as possible, but without colliding
with each other. We aim to minimize the maximum time any agent takes to reach
their goal, ensuring optimal path length. In this work, we complement a recent
thread of results that aim to systematically study the algorithmic behavior of
this problem, through the parameterized complexity point of view.
  First, we show that MAPF is NP-hard when the given network has a star-like
topology (bounded vertex cover number) or is a tree with $11$ leaves. Both of
these results fill important gaps in our understanding of the tractability of
this problem that were left untreated in the recent work of [Fioravantes et al.
Exact Algorithms and Lowerbounds for Multiagent Path Finding: Power of Treelike
Topology. AAAI'24]. Nevertheless, our main contribution is an exact algorithm
that scales well as the input grows (FPT) when the topology of the given
network is highly centralized (bounded distance to clique). This parameter is
significant as it mirrors real-world networks. In such environments, a bunch of
central hubs (e.g., processing areas) are connected to only few peripheral
nodes.


## AI 摘要

多智能体路径规划(MAPF)问题研究如何为多个智能体规划无碰撞的最优路径。最新研究从参数化复杂度角度分析了该问题：当网络呈星型拓扑(有界顶点覆盖数)或11叶树结构时，MAPF是NP难问题。研究主要贡献是提出了一种适用于高度中心化网络(有界团距离)的精确算法，该参数能很好模拟现实网络(如中心枢纽连接少量外围节点的场景)。算法在输入规模增大时仍保持良好扩展性(FPT)，填补了前人研究中关于树状拓扑可解性的重要空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T20:02:29Z
- **目录日期**: 2025-06-03
