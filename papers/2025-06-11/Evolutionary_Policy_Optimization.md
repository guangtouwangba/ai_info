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

本文提出了一种混合算法——进化策略优化（EPO），结合了进化算法（EA）的可扩展性和多样性，以及策略梯度方法的性能和稳定性。EPO通过维护基于潜在变量的智能体种群，共享网络参数以提高效率和一致性，并将多样化经验聚合到主智能体中。实验表明，在灵巧操作、腿部运动和经典控制任务中，EPO在样本效率、渐进性能和可扩展性方面均优于现有最先进方法。该算法解决了策略梯度算法在大批量训练时的冗余问题，同时弥补了进化算法样本效率低的不足。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T20:02:07Z
- **目录日期**: 2025-06-11
