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

本文提出了一种静默、自稳定的分布式排名协议，适用于基于随机配对交互的群体协议模型。该协议为n个匿名代理分配唯一排名（1到n），同时也可用于自稳定的领导者选举。核心创新在于仅需n + O(log n)状态的空间高效非自稳定排名协议。最终协议在O(n² log n)交互次数内以高概率稳定，总状态数为n + O(log² n)，显著优于现有最佳协议（需n + Ω(n)状态）。该结果在状态开销上实现了指数级改进，且稳定时间达到渐近最优。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T13:09:15Z
- **目录日期**: 2025-04-15
