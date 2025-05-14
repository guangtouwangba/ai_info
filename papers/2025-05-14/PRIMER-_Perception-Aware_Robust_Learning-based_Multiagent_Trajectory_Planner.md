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

本文提出了两种感知感知的分散式异步多智能体轨迹规划器PARM和PARM*，用于在定位不确定环境下实现避碰导航，其中PARM*通过更多计算获得更优解。虽然性能优越，但计算成本较高。为此，作者进一步提出基于模仿学习的神经网络规划器PRIMER，以PARM*为专家示范进行训练。PRIMER在部署时利用神经网络的低计算需求，计算速度比优化方法快达5500倍，解决了实时重规划难题。该研究通过结合优化方法和学习技术，实现了高效可靠的分散式多智能体轨迹规划。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T10:02:20Z
- **目录日期**: 2025-05-14
