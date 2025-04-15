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

本文提出了一种静默自稳定的分布式排名协议，用于解决群体协议模型中的领导者选举问题。该协议能够在$O(n^2\log n)$次交互内以高概率稳定，仅需$n + O(\log^2 n)$个状态空间，相比现有最优方案将额外状态数从$Ω(n)$指数级降低到$O(\log^2 n)$。协议核心是一个空间高效的（非自稳定）基础排名算法，仅需$n + O(\log n)$状态。该成果在状态空间和稳定时间（达到理论下界）方面均取得显著优化，同时实现了自稳定的领导者选举功能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T14:02:07Z
- **目录日期**: 2025-04-15
