# A Dataless Reinforcement Learning Approach to Rounding Hyperplane Optimization for Max-Cut

**URL**: http://arxiv.org/abs/2505.13405v2

## 原始摘要

The Maximum Cut (MaxCut) problem is NP-Complete, and obtaining its optimal
solution is NP-hard in the worst case. As a result, heuristic-based algorithms
are commonly used, though their design often requires significant domain
expertise. More recently, learning-based methods trained on large (un)labeled
datasets have been proposed; however, these approaches often struggle with
generalizability and scalability. A well-known approximation algorithm for
MaxCut is the Goemans-Williamson (GW) algorithm, which relaxes the Quadratic
Unconstrained Binary Optimization (QUBO) formulation into a semidefinite
program (SDP). The GW algorithm then applies hyperplane rounding by uniformly
sampling a random hyperplane to convert the SDP solution into binary node
assignments. In this paper, we propose a training-data-free approach based on a
non-episodic reinforcement learning formulation, in which an agent learns to
select improved rounding hyperplanes that yield better cuts than those produced
by the GW algorithm. By optimizing over a Markov Decision Process (MDP), our
method consistently achieves better cuts across large-scale graphs with varying
densities and degree distributions.


## AI 摘要

本文提出了一种基于强化学习的无训练数据方法，用于改进最大割问题(MaxCut)的Goemans-Williamson(GW)算法。传统GW算法通过随机采样超平面进行舍入，而该方法通过马尔可夫决策过程(MDP)学习选择更优的舍入超平面。与需要大量训练数据的学习方法不同，这种无训练数据方法能适应不同密度和度分布的大规模图，始终获得比GW算法更好的割结果。该方法克服了现有启发式算法需要专业知识以及学习算法泛化性和可扩展性不足的问题，为NP难的最大割问题提供了新的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T02:30:25Z
- **目录日期**: 2025-05-21
