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

多智能体路径规划(MAPF)问题研究如何为多个智能体规划无碰撞的最优路径。研究表明，在星型拓扑网络或具有11个叶子的树状网络中，该问题是NP难解的。作者提出了一种针对高度中心化网络(如具有有限距离到完全图的网络)的精确算法，该算法在参数化复杂度框架下表现良好(FPT)。这类网络结构(如具有中心枢纽和少量外围节点的处理区域)在现实场景中很常见。该研究填补了前人工作中关于树状拓扑结构可解性的空白，为实际应用提供了理论基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T02:32:30Z
- **目录日期**: 2025-06-04
