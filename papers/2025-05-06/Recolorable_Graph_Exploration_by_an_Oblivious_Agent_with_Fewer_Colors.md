# Recolorable Graph Exploration by an Oblivious Agent with Fewer Colors

**URL**: http://arxiv.org/abs/2505.02789v1

## 原始摘要

Recently, B\"ockenhauer, Frei, Unger, and Wehner (SIROCCO 2023) introduced a
novel variant of the graph exploration problem in which a single memoryless
agent must visit all nodes of an unknown, undirected, and connected graph
before returning to its starting node. Unlike the standard model for mobile
agents, edges are not labeled with port numbers. Instead, the agent can color
its current node and observe the color of each neighboring node. To move, it
specifies a target color and then moves to an adversarially chosen neighbor of
that color. B\"ockenhauer~et al.~analyzed the minimum number of colors required
for successful exploration and proposed an elegant algorithm that enables the
agent to explore an arbitrary graph using only eight colors. In this paper, we
present a novel graph exploration algorithm that requires only six colors.
Furthermore, we prove that five colors are sufficient if we consider only a
restricted class of graphs, which we call the $\varphi$-free graphs, a class
that includes every graph with maximum degree at most three and every cactus.


## AI 摘要

本文研究了一种新型图探索问题，其中无记忆代理需在不含端口号的未知无向连通图中遍历所有节点并返回起点。代理通过着色当前节点并观察相邻节点颜色进行导航，通过指定目标颜色移动。Böckenhauer等人(2023)提出使用8种颜色的算法，本文则提出仅需6种颜色的新算法。此外，证明在φ-free图类(包括最大度≤3的图和仙人掌图)中，5种颜色即足够。该研究改进了原有结果，降低了探索所需颜色数。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T04:01:55Z
- **目录日期**: 2025-05-06
