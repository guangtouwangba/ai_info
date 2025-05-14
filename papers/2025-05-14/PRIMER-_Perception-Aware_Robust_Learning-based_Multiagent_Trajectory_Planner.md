# PRIMER: Perception-Aware Robust Learning-based Multiagent Trajectory Planner

**URL**: http://arxiv.org/abs/2406.10060v4

## 原始摘要

In decentralized multiagent trajectory planners, agents need to communicate
and exchange their positions to generate collision-free trajectories. However,
due to localization errors/uncertainties, trajectory deconfliction can fail
even if trajectories are perfectly shared between agents. To address this
issue, we first present PARM and PARM*, perception-aware, decentralized,
asynchronous multiagent trajectory planners that enable a team of agents to
navigate uncertain environments while deconflicting trajectories and avoiding
obstacles using perception information. PARM* differs from PARM as it is less
conservative, using more computation to find closer-to-optimal solutions. While
these methods achieve state-of-the-art performance, they suffer from high
computational costs as they need to solve large optimization problems onboard,
making it difficult for agents to replan at high rates. To overcome this
challenge, we present our second key contribution, PRIMER, a learning-based
planner trained with imitation learning (IL) using PARM* as the expert
demonstrator. PRIMER leverages the low computational requirements at deployment
of neural networks and achieves a computation speed up to 5500 times faster
than optimization-based approaches.


## AI 摘要

本文提出了两种感知感知的分散式异步多智能体轨迹规划器PARM和PARM*，用于在不确定环境中导航并避免碰撞。PARM*比PARM更优化但计算成本较高。为降低计算负担，作者进一步提出基于模仿学习的PRIMER规划器，以PARM*为专家示范进行训练。PRIMER利用神经网络在部署时的低计算需求，实现了比优化方法快5500倍的计算速度，同时保持高性能。这些方法解决了定位误差导致的轨迹冲突问题，并通过感知信息实现避障。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T06:02:10Z
- **目录日期**: 2025-05-14
