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

这篇论文提出了一种改进的图探索算法，使无记忆智能体仅需6种颜色（而非之前的8种）即可探索未知无向连通图。在智能体无法通过端口号识别边的模型中，该算法通过着色当前节点并观察邻接节点颜色来导航。此外，研究证明对于特定φ-free图类（包括最大度≤3的图和仙人掌图），5种颜色即足够。该成果改进了Böckenhauer等人(SIROCCO 2023)的工作，降低了完成图探索任务所需的颜色复杂度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T01:29:35Z
- **目录日期**: 2025-05-07
