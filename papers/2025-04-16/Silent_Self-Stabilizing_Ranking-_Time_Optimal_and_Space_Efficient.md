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

本文提出了一种静默、自稳定的分布式排名协议，适用于随机配对交互的群体计算模型。该协议为n个匿名代理分配唯一排名{1,...,n}，同时可作为自稳定的领导者选举协议。核心创新在于仅需n+O(log n)状态的高效非自稳排名协议，改进后版本使用n+O(log²n)状态，显著优于现有需要n+Ω(n)状态的方案。协议以O(n²logn)交互次数稳定（概率高且期望值），达到渐进最优稳定时间。相比现有最佳方案，本协议将额外状态数从线性降低至对数级，实现指数级改进。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T00:02:37Z
- **目录日期**: 2025-04-16
