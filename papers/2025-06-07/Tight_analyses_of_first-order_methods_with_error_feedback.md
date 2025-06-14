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

本文分析了分布式学习中两种常见的误差反馈通信压缩方法（EF和EF21），通过构建最优Lyapunov函数，为两种方法提供了紧收敛率分析（含匹配下界）。研究采用简化但具代表性的设定，不仅获得清晰理论见解，还实现了EF、EF21与压缩梯度下降法的公平比较。结果表明，该方法能有效缓解通信压缩导致的收敛性能下降问题，为通信瓶颈优化提供了理论依据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T23:02:27Z
- **目录日期**: 2025-06-07
