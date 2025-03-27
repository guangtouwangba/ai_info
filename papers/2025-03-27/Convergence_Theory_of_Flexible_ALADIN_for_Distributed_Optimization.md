# Convergence Theory of Flexible ALADIN for Distributed Optimization

**URL**: http://arxiv.org/abs/2503.20716v1

## 原始摘要

The Augmented Lagrangian Alternating Direction Inexact Newton (ALADIN) method
is a cutting-edge distributed optimization algorithm known for its superior
numerical performance. It relies on each agent transmitting information to a
central coordinator for data exchange. However, in practical network
optimization and federated learning, unreliable information transmission often
leads to packet loss, posing challenges for the convergence analysis of ALADIN.
To address this issue, this paper proposes Flexible ALADIN, a random polling
variant of ALADIN, and presents a rigorous convergence analysis, including
global convergence for convex problems and local convergence for non-convex
problems.


## AI 摘要

本文提出了Flexible ALADIN算法，这是ALADIN分布式优化算法的随机轮询改进版本，旨在解决实际应用中因信息传输不可靠导致的数据包丢失问题。该研究提供了严格的收敛性分析，证明了其在凸问题上的全局收敛性以及在非凸问题上的局部收敛性。这一方法特别适用于网络优化和联邦学习等场景，通过改进信息交换机制，增强了算法在不可靠通信环境下的鲁棒性，同时保持了ALADIN原有的优异数值性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T20:01:47Z
- **目录日期**: 2025-03-27
