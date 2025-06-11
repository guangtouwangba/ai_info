# Evolutionary Policy Optimization

**URL**: http://arxiv.org/abs/2503.19037v2

## 原始摘要

On-policy reinforcement learning (RL) algorithms are widely used for their
strong asymptotic performance and training stability, but they struggle to
scale with larger batch sizes, as additional parallel environments yield
redundant data due to limited policy-induced diversity. In contrast,
Evolutionary Algorithms (EAs) scale naturally and encourage exploration via
randomized population-based search, but are often sample-inefficient. We
propose Evolutionary Policy Optimization (EPO), a hybrid algorithm that
combines the scalability and diversity of EAs with the performance and
stability of policy gradients. EPO maintains a population of agents conditioned
on latent variables, shares actor-critic network parameters for coherence and
memory efficiency, and aggregates diverse experiences into a master agent.
Across tasks in dexterous manipulation, legged locomotion, and classic control,
EPO outperforms state-of-the-art baselines in sample efficiency, asymptotic
performance, and scalability.


## AI 摘要

本文提出了一种结合进化算法（EA）和策略梯度优点的混合算法——进化策略优化（EPO）。传统策略梯度方法在大批量训练时因策略多样性不足导致数据冗余，而进化算法虽具扩展性和探索性，但样本效率较低。EPO通过维护基于隐变量的智能体种群，共享网络参数以提高一致性和内存效率，并将多样化经验聚合到主智能体中。实验表明，EPO在灵巧操作、腿部运动及经典控制任务中，在样本效率、渐进性能和可扩展性上均优于现有最优方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T12:02:22Z
- **目录日期**: 2025-06-11
