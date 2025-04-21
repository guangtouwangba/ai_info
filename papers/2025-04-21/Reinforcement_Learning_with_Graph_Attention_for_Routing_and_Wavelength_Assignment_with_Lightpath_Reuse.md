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

本文研究了固定网格光网络中支持灵活速率收发器的路由与波长分配问题（RWA-LR），通过重用现有光路来提升服务容量。研究发现，按跳数排序候选路径的启发式算法比按总长度排序提升6%吞吐量。作者采用图注意力网络构建强化学习（RL）智能体，其策略优于现有最佳RL方法2.5%（17.4Tbps）和最佳启发式算法1.2%（8.5Tbps）。尽管边际增益有限，但凸显了在长期资源分配任务中学习有效RL策略的挑战。研究方法和代码均已开源。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T23:02:10Z
- **目录日期**: 2025-04-21
