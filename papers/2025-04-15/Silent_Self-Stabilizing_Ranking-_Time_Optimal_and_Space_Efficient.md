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

本文提出了一种基于群体协议模型的静默自稳定排序协议，用于为n个匿名代理分配唯一排名（1到n）。该协议同时可作为自稳定领导者选举协议。核心是一个仅需n+O(log n)状态的高效排序算法，最终协议使用n+O(log² n)状态，显著优于现有需要n+Ω(n)状态的方案。协议以O(n² log n)交互次数稳定，该时间复杂度是渐进最优的（与Burman等人PODC'21结果相比）。本研究在状态空间复杂度上实现了指数级改进，将额外状态从线性量级降至对数平方量级。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T07:01:58Z
- **目录日期**: 2025-04-15
