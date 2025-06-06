# Tight analyses of first-order methods with error feedback

**URL**: http://arxiv.org/abs/2506.05271v1

## 原始摘要

Communication between agents often constitutes a major computational
bottleneck in distributed learning. One of the most common mitigation
strategies is to compress the information exchanged, thereby reducing
communication overhead. To counteract the degradation in convergence associated
with compressed communication, error feedback schemes -- most notably
$\mathrm{EF}$ and $\mathrm{EF}^{21}$ -- were introduced. In this work, we
provide a tight analysis of both of these methods. Specifically, we find the
Lyapunov function that yields the best possible convergence rate for each
method -- with matching lower bounds. This principled approach yields sharp
performance guarantees and enables a rigorous, apples-to-apples comparison
between $\mathrm{EF}$, $\mathrm{EF}^{21}$, and compressed gradient descent. Our
analysis is carried out in a simplified yet representative setting, which
allows for clean theoretical insights and fair comparison of the underlying
mechanisms.


## AI 摘要

本文分析了分布式学习中通信压缩带来的收敛性下降问题，重点研究了两种误差反馈方案（EF和EF21）。通过构建最优Lyapunov函数，研究获得了两种方法的最佳收敛率及其匹配下界，提供了严格的性能保证。这种系统性方法实现了EF、EF21和压缩梯度下降之间的公平比较。研究在简化但具代表性的设置下进行，既保证了理论清晰性，又能准确评估不同机制的核心性能。结果表明，误差反馈能有效缓解通信压缩导致的收敛退化问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T09:02:27Z
- **目录日期**: 2025-06-06
