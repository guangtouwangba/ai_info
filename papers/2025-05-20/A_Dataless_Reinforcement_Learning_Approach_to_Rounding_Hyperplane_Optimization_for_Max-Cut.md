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

本文提出了一种基于强化学习的无训练数据方法，用于改进最大割问题(MaxCut)的近似解。传统Goemans-Williamson(GW)算法通过随机超平面划分获得近似解，而新方法通过马尔可夫决策过程(MDP)学习选择更优的划分超平面。与需要大量标注数据的学习方法不同，该方法不依赖训练数据，能适应不同密度和度分布的大规模图，始终获得优于GW算法的切割结果。该方法解决了现有学习型方法在泛化性和可扩展性方面的不足，同时避免了启发式算法对领域专业知识的依赖。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T22:01:11Z
- **目录日期**: 2025-05-20
