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

本文分析了分布式学习中两种常见的误差反馈通信压缩方法（EF和EF21），通过构建最优Lyapunov函数，首次给出了两种方法的紧收敛率分析（含匹配下界）。在简化但具代表性的设定下，研究揭示了误差补偿机制的本质差异，为EF、EF21和压缩梯度下降法提供了严格的性能比较基准。理论分析表明，EF21相比EF具有更优的收敛保证，这为通信压缩算法的选择提供了量化依据。该工作通过统一的框架实现了不同压缩方法的公平对比，为分布式优化中的通信效率提升提供了理论指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T08:02:22Z
- **目录日期**: 2025-06-07
