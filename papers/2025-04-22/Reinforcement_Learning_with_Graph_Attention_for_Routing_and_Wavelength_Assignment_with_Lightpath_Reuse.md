# Reinforcement Learning with Graph Attention for Routing and Wavelength Assignment with Lightpath Reuse

**URL**: http://arxiv.org/abs/2502.14741v2

## 原始摘要

Many works have investigated reinforcement learning (RL) for routing and
spectrum assignment on flex-grid networks but only one work to date has
examined RL for fixed-grid with flex-rate transponders, despite production
systems using this paradigm. Flex-rate transponders allow existing lightpaths
to accommodate new services, a task we term routing and wavelength assignment
with lightpath reuse (RWA-LR). We re-examine this problem and present a
thorough benchmarking of heuristic algorithms for RWA-LR, which are shown to
have 6% increased throughput when candidate paths are ordered by number of
hops, rather than total length. We train an RL agent for RWA-LR with graph
attention networks for the policy and value functions to exploit the
graph-structured data. We provide details of our methodology and open source
all of our code for reproduction. We outperform the previous state-of-the-art
RL approach by 2.5% (17.4 Tbps mean additional throughput) and the best
heuristic by 1.2% (8.5 Tbps mean additional throughput). This marginal gain
highlights the difficulty in learning effective RL policies on long horizon
resource allocation tasks.


## AI 摘要

该研究重新审视了固定网格网络中带光路复用的路由与波长分配问题(RWA-LR)，提出了一种基于图注意力网络的强化学习方法。相比传统启发式算法（按跳数排序路径可提升6%吞吐量），该方法以17.4Tbps的额外吞吐量超越现有最佳强化学习方法2.5%，并以8.5Tbps优势超越最佳启发式算法1.2%。研究开源了全部代码，并指出在长周期资源分配任务中学习有效强化学习策略的困难性。成果突显了图神经网络在处理图结构数据上的优势，为实际生产系统提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T01:29:40Z
- **目录日期**: 2025-04-22
