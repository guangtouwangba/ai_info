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

本文提出了一种无需训练数据的强化学习方法来解决最大割问题(MaxCut)。该方法基于Goemans-Williamson(GW)算法框架，通过马尔可夫决策过程(MDP)学习选择更优的切割超平面，而非像GW算法那样随机采样。相比传统启发式算法和现有基于学习的方法，这种强化学习方案能在大规模不同密度和度分布的图上持续获得更好的切割结果，同时避免了数据依赖性和可扩展性问题。该方法创新性地将GW算法的半定规划松弛与强化学习的决策优化相结合，为NP难的最大割问题提供了新的解决思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T03:19:05Z
- **目录日期**: 2025-05-20
