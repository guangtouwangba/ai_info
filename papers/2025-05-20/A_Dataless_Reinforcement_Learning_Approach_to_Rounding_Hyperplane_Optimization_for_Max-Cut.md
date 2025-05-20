# A Dataless Reinforcement Learning Approach to Rounding Hyperplane Optimization for Max-Cut

**URL**: http://arxiv.org/abs/2505.13405v1

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

本文提出了一种基于强化学习的无训练数据方法，用于改进最大割问题的近似解。传统Goemans-Williamson(GW)算法通过随机超平面分割求解半定规划松弛，而本方法通过马尔可夫决策过程训练智能体学习选择更优的分割超平面。相比需要大量标注数据的学习方法，该方案具有更好的泛化性和可扩展性。实验表明，该方法在不同密度和度分布的大规模图上，能持续获得比GW算法更优的切割结果。这种免训练的特性使其避免了传统学习方法的局限性，同时超越了经典近似算法的性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T09:01:28Z
- **目录日期**: 2025-05-20
