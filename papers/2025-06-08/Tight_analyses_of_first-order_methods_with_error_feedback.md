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

本文分析了分布式学习中通信压缩导致的收敛性下降问题，重点研究了两种误差反馈方案（EF和EF21）。通过构建最优Lyapunov函数，作者获得了两种方法的最紧收敛率，并给出了匹配的下界。这种理论分析方法不仅提供了精确的性能保证，还实现了EF、EF21和压缩梯度下降之间的公平比较。研究在一个简化但具有代表性的设置下进行，既保证了理论清晰度，又能有效比较不同机制的核心特性。结果表明该方法可为通信压缩策略的选择提供严格的理论依据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-08T09:02:18Z
- **目录日期**: 2025-06-08
