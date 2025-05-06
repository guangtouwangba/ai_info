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

本文提出了一种改进的图探索算法，将所需颜色数从之前的8种减少到6种。研究针对无记忆代理在无端口号标记的未知无向连通图中探索的问题，代理通过着色节点并观察邻接节点颜色来导航。对于特定类型的φ-free图（包括最大度不超过3的图和仙人掌图），证明5种颜色即足够。该成果改进了Böckenhauer等人先前的工作，在更广泛图类上实现了更优的探索效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T09:01:54Z
- **目录日期**: 2025-05-06
