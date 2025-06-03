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

多智能体路径规划（MAPF）问题旨在为多个智能体在给定网络中规划无碰撞的最优路径，以最小化最大到达时间。研究表明，该问题在星型拓扑（有限顶点覆盖数）或11叶树状网络中属于NP难问题，填补了先前研究的空白。主要贡献是提出了一种针对高度集中化网络（有限距离团结构）的精确算法，其计算复杂度随输入规模增长可控（FPT参数化算法）。这类网络结构（如中心枢纽连接少量外围节点）在现实场景（如处理区域）中具有重要应用价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T13:12:32Z
- **目录日期**: 2025-06-03
