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

该研究重新审视了固定网格网络中带可复用光路径的路由和波长分配问题(RWA-LR)，对启发式算法进行了全面基准测试，发现按跳数排序路径比按总长度排序能提升6%吞吐量。研究者采用图注意力网络构建强化学习(RL)代理的策略和价值函数，利用图结构数据。相比之前最优RL方法提升2.5%(17.4Tbps)，比最佳启发式算法提升1.2%(8.5Tbps)。这些边际增益表明，在长周期资源分配任务中学习有效的RL策略具有挑战性。研究开源了所有代码以促进复现。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T14:02:32Z
- **目录日期**: 2025-04-21
