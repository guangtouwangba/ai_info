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

本文研究了多智能体路径规划(MAPF)问题，证明在星型拓扑网络(顶点覆盖数有限)和11叶树结构下该问题是NP难的，填补了前人研究的空白。主要贡献是提出了一种针对高度中心化网络(有限团距离)的精确算法，该算法在输入规模增大时(FPT)仍能保持良好性能。这种参数化方法具有现实意义，因为许多实际网络(如处理中心连接少量外围节点)都具有这种中心化拓扑结构。研究从参数化复杂度角度系统分析了MAPF问题的算法行为，为理解该问题的可解性提供了新见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T10:02:23Z
- **目录日期**: 2025-06-03
