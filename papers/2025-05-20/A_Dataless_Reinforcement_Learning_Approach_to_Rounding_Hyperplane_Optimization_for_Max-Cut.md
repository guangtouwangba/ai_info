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

本文提出了一种基于无训练数据的强化学习方法来解决最大割问题(MaxCut)。该方法改进了经典的Goemans-Williamson(GW)算法中的超平面舍入步骤，通过马尔可夫决策过程(MDP)学习选择更优的舍入超平面，而非GW算法中使用的随机超平面。相比需要大量训练数据的现有学习方法，这种无训练数据的方法在大规模不同密度和度分布的图上都能稳定获得比GW算法更好的切割结果，同时避免了传统启发式算法需要专业知识的问题，也克服了学习型方法在泛化性和可扩展性方面的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T20:01:04Z
- **目录日期**: 2025-05-20
