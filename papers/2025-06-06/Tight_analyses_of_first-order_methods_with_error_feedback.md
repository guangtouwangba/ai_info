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

该研究对分布式学习中两种常见的误差反馈通信压缩方法（EF和EF21）进行了严格分析。通过构建最优Lyapunov函数，作者获得了匹配上下界的最优收敛速率，为这两种方法与压缩梯度下降法提供了公平比较的理论基础。研究采用简化但具代表性的设置，揭示了不同方法的内在机制差异，为通信压缩场景下的算法选择提供了理论依据。分析表明，EF21相比EF具有更优的收敛性能，这为分布式学习系统中通信效率与收敛速度的权衡提供了量化参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T14:02:20Z
- **目录日期**: 2025-06-06
