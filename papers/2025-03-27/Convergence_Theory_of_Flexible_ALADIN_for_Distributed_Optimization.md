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

本文提出了Flexible ALADIN算法，这是对现有ALADIN分布式优化方法的改进版本。针对实际应用中信息传输不可靠导致丢包的问题，该算法采用随机轮询机制，并提供了严格的收敛性分析：在凸问题上具有全局收敛性，在非凸问题上具有局部收敛性。这一改进使ALADIN方法在存在通信不可靠的网络优化和联邦学习场景中更具实用性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T17:01:58Z
- **目录日期**: 2025-03-27
