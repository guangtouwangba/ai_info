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

这篇论文改进了Böckenhauer等人提出的图探索问题模型，其中无记忆智能体需在未知无向连通图中遍历所有节点后返回起点。新模型不依赖边端口号，而是通过节点着色进行导航。原算法需8种颜色，本文提出仅需6色的新算法。此外，证明在φ-free图类（包含所有最大度≤3的图和仙人掌图）中5色即足够。研究显著降低了所需颜色数量，为特定图类提供了更优解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T22:01:48Z
- **目录日期**: 2025-05-06
