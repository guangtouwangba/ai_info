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

分布式学习中，通信常成为计算瓶颈，通常采用信息压缩来降低通信开销。为缓解压缩通信导致的收敛性能下降，误差反馈方法（如EF和EF21）被提出。本研究对这两种方法进行了精确分析，通过寻找最优Lyapunov函数确定了各自的最佳收敛速率，并给出了匹配下界。这种系统性方法提供了严格的性能保证，使EF、EF21与压缩梯度下降法能在统一框架下公平比较。分析基于简化但具代表性的设定，既保证了理论清晰度，又能揭示不同机制的本质差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T06:02:24Z
- **目录日期**: 2025-06-07
