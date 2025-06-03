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

多智能体路径规划(MAPF)问题研究如何为多个智能体规划无碰撞的最优路径。最新研究表明：在星型拓扑网络或11叶树结构下，该问题是NP难解的。针对实际中常见的中心化网络拓扑(如带有限制距离的团结构)，研究者提出了一个高效的固定参数可解(FPT)算法。该算法特别适用于具有中心枢纽(如处理区域)和少量外围节点的现实网络场景，能随着输入规模增长保持良好扩展性。这些成果填补了前人研究空白，为理解MAPF问题的计算复杂度提供了新见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T09:02:36Z
- **目录日期**: 2025-06-03
