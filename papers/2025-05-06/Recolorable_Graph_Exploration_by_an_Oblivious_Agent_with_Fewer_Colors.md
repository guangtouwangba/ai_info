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

本文提出了一种改进的图探索算法，使无记忆智能体仅需6种颜色即可探索未知无向连通图，优于先前需要8种颜色的算法。对于特定类型的φ-free图（包括最大度数≤3的图和仙人掌图），证明5种颜色足够。该模型不依赖边端口号，而是通过节点着色和邻域颜色观察进行导航。智能体通过指定目标颜色移动，对手选择对应颜色的邻居。研究改进了Böckenhauer等人(SIROCCO 2023)的工作，在更广的图类上降低了颜色需求，为图探索问题提供了更优解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T15:02:43Z
- **目录日期**: 2025-05-06
