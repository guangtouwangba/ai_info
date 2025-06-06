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

分布式学习中，代理间的通信常成为计算瓶颈。为减少通信开销，压缩信息交换是常见策略，但会降低收敛性。误差反馈方法（如EF和EF21）被提出以缓解该问题。本研究对这两种方法进行了精确分析，通过寻找最优Lyapunov函数，得出了各自的最佳收敛率及匹配下界。这种系统性方法提供了清晰的性能保证，并支持EF、EF21与压缩梯度下降法的公平比较。分析在简化但具代表性的设置下进行，有助于获得干净的理论见解和机制对比。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T18:02:15Z
- **目录日期**: 2025-06-06
