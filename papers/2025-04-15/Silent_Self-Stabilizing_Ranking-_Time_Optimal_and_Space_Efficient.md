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

本文提出了一种静默自稳定的分布式排名协议，用于解决种群协议模型中的领导者选举问题。该协议通过随机配对交互，为n个匿名代理分配唯一排名（1到n），从而间接实现领导者选举。核心创新在于仅需n+O(log²n)状态空间（相比现有最佳方案指数级减少）和O(n²logn)交互次数（理论最优）。协议结合了高效非自稳定排名算法（n+O(logn)状态）与自稳定机制，显著降低了额外状态开销，同时保证了概率高和期望下的快速收敛。这一成果在状态空间效率上实现了突破性进展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T17:02:35Z
- **目录日期**: 2025-04-15
