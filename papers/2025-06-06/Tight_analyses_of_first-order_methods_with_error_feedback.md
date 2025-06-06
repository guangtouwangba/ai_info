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

该研究分析了分布式学习中两种常见的误差反馈通信压缩方法（EF和EF21），通过构建最优Lyapunov函数，首次给出了两种方法的紧收敛率分析（含匹配下界）。在简化但具代表性的实验设置下，研究不仅提供了精确的性能保证，还实现了EF、EF21与压缩梯度下降法的公平比较。这项工作揭示了误差反馈机制对抗通信压缩导致收敛性能下降的本质原理，为通信受限场景下的算法选择提供了理论依据。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T07:02:36Z
- **目录日期**: 2025-06-06
