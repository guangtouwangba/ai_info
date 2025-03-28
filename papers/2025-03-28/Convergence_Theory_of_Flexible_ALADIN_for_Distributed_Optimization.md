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

本文提出了一种改进的分布式优化算法——柔性ALADIN（Flexible ALADIN），以解决传统ALADIN方法在信息传输不可靠时（如丢包）的收敛性问题。该算法采用随机轮询机制，适用于网络优化和联邦学习等场景。研究提供了严格的收敛性分析，证明了其在凸问题上的全局收敛性，以及在非凸问题上的局部收敛性。这一创新既保持了ALADIN优异的数值性能，又增强了其在现实通信环境中的鲁棒性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T02:28:25Z
- **目录日期**: 2025-03-28
