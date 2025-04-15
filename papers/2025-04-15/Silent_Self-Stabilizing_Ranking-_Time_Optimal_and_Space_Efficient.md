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

本文提出了一种静默自稳定的分布式排序协议，用于在随机配对的代理交互模型中为n个匿名代理分配唯一排名（1到n）。该协议也可用于自稳定的领导者选举。核心创新是一个仅需n+O(log n)状态的高效非自稳定排序算法，改进为自稳定版本后总状态数为n+O(log² n)，显著优于现有需要n+Ω(n)状态的方案。协议以O(n² log n)交互次数稳定（概率高且期望值），该时间复杂度已被证明是最优的。相比现有最佳方案，本协议将额外状态数从线性降低到对数级，实现了指数级改进。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T18:02:28Z
- **目录日期**: 2025-04-15
