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

本文提出了一种静默自稳定的分布式排名协议，用于解决匿名代理间的唯一排序问题。该协议采用两两随机交互的群体协议模型，核心是一个仅需n+O(log n)状态的高效非自稳定排序算法。最终协议在O(n² log n)次交互内以高概率稳定，总状态数为n+O(log² n)，达到渐进最优的稳定时间。相比现有最佳方案所需的n+Ω(n)状态，本协议将额外状态数指数级降低。该协议同时实现了自稳定的领导者选举功能，为分布式系统提供了一种更高效的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T08:01:56Z
- **目录日期**: 2025-04-15
