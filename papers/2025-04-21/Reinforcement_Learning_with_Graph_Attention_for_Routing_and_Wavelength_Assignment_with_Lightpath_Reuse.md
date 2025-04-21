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

本文研究了固定网格网络中带可调谐速率收发器的路由与波长分配问题（RWA-LR）。相比传统方法，按跳数排序候选路径可使吞吐量提升6%。作者采用图注意力网络构建强化学习（RL）智能体来处理图结构数据，其性能比现有最佳RL方法提升2.5%（17.4Tbps），比最优启发式算法提升1.2%（8.5Tbps）。研究开源了全部代码，并指出在长周期资源分配任务中学习有效RL策略的挑战性。结果表明，尽管性能提升有限，但RL在光网络优化中仍具潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T04:02:28Z
- **目录日期**: 2025-04-21
