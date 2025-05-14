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

本文提出了两种感知感知的分散式多智能体轨迹规划器PARM和PARM*，能处理定位不确定性下的轨迹避碰问题，其中PARM*计算量更大但解更优。为克服这些优化方法计算量大的缺点，作者进一步提出基于模仿学习的神经网络规划器PRIMER，以PARM*为专家示范进行训练。PRIMER在保持性能的同时，部署时计算速度比优化方法快达5500倍，实现了实时高性能的多智能体轨迹规划。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T17:02:42Z
- **目录日期**: 2025-05-14
