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

该研究提出了两种去中心化多智能体轨迹规划器PARM和PARM*，通过感知信息实现不确定环境下的避障与轨迹协调。PARM*相比PARM计算量更大但解更优。为克服优化方法计算成本高的问题，研究者进一步开发了基于模仿学习的神经网络规划器PRIMER，以PARM*为专家示范进行训练。PRIMER在部署时计算速度比优化方法快达5500倍，同时保持优异性能。这些方法解决了定位误差导致的轨迹冲突问题，显著提升了多智能体系统的实时规划能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T21:02:08Z
- **目录日期**: 2025-05-14
