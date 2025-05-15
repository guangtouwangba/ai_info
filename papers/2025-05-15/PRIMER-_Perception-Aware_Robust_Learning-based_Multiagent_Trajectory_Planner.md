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

本文提出两种感知感知的分散式多智能体轨迹规划器PARM和PARM*，通过利用感知信息解决定位不确定性导致的轨迹冲突问题，其中PARM*计算更密集但解更优。虽然性能先进，但优化计算成本高。为此，作者进一步提出基于模仿学习的神经网络规划器PRIMER，以PARM*为专家示范进行训练，在保持性能的同时将计算速度提升高达5500倍。PRIMER结合了优化方法的性能优势和神经网络的部署效率，实现了高速重规划能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T01:30:16Z
- **目录日期**: 2025-05-15
