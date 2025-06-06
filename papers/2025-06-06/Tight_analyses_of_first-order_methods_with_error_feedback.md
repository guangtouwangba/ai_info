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

分布式学习中，智能体间的通信常成为计算瓶颈。为减少通信开销，通常采用信息压缩策略，但这会降低收敛性能。为此，误差反馈机制（如EF和EF21）被提出。本研究对这两种方法进行了精确分析，通过构建最优Lyapunov函数确定了各自的最佳收敛速率（含匹配下界），从而提供严格的性能保证，并实现EF、EF21与压缩梯度下降的公平比较。该分析在简化但具代表性的设置下进行，能清晰揭示理论机制并支持客观对比。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T19:02:13Z
- **目录日期**: 2025-06-06
