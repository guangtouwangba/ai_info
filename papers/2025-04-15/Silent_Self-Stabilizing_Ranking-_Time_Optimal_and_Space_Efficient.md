# Silent Self-Stabilizing Ranking: Time Optimal and Space Efficient

**URL**: http://arxiv.org/abs/2504.10417v1

## 原始摘要

We present a silent, self-stabilizing ranking protocol for the population
protocol model of distributed computing, where agents interact in randomly
chosen pairs to solve a common task. We are given $n$ anonymous agents, and the
goal is to assign each agent a unique rank in $\{1, \dots, n\}$. Given unique
ranks, it is straightforward to select a designated leader. Thus, our protocol
is a self-stabilizing leader election protocol as well. Ranking requires at
least $n$ states per agent; hence, the goal is to minimize the additional
number of states, called overhead states. The core of our protocol is a
space-efficient but non-self-stabilizing ranking protocol that requires only $n
+ O(\log n)$ states. Our protocol stabilizes in $O(n^2\log n)$ interactions
w.h.p.\ and in expectation, using $n + O(\log^2 n)$ states in total. Our
stabilization time is asymptotically optimal (see Burman et al., PODC'21). In
comparison to the currently best known ranking protocol by Burman et al., which
requires $n + \Omega(n)$ states, our result exponentially improves the number
of overhead states.


## AI 摘要

本文提出了一种静默自稳定的分布式排名协议，用于解决匿名代理间的唯一排名问题。该协议通过两两随机交互实现，仅需n + O(log²n)个状态，显著优于现有需要n + Ω(n)状态的方案。协议核心是一个空间高效的（非自稳定）排名算法，仅需n + O(log n)状态。理论分析表明，该协议在O(n²log n)次交互后以高概率稳定，且稳定时间达到渐进最优。作为副产品，该协议也可实现自稳定的领导者选举。相比现有最佳方案，本工作实现了状态数量的指数级优化。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T19:02:26Z
- **目录日期**: 2025-04-15
