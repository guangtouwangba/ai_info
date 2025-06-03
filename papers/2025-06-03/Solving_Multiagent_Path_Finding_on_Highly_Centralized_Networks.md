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

本文研究了多智能体路径规划(MAPF)问题的计算复杂性。首先证明了在星型网络(顶点覆盖数有限)或11个叶子的树状拓扑中,MAPF是NP难问题。主要贡献是提出了一种精确算法，适用于高度中心化的网络拓扑(有限距离团数)，这类参数化算法能随着输入规模良好扩展(FPT)。这种网络结构在现实中很常见，表现为少数中心枢纽(如处理区域)连接多个外围节点。研究填补了前人工作中关于树状拓扑可解性的空白，为实际应用提供了理论基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T12:02:21Z
- **目录日期**: 2025-06-03
