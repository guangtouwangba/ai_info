# Fair Dynamic Spectrum Access via Fully Decentralized Multi-Agent Reinforcement Learning

**URL**: http://arxiv.org/abs/2503.24296v1

## 原始摘要

We consider a decentralized wireless network with several source-destination
pairs sharing a limited number of orthogonal frequency bands. Sources learn to
adapt their transmissions (specifically, their band selection strategy) over
time, in a decentralized manner, without sharing information with each other.
Sources can only observe the outcome of their own transmissions (i.e., success
or collision), having no prior knowledge of the network size or of the
transmission strategy of other sources. The goal of each source is to maximize
their own throughput while striving for network-wide fairness. We propose a
novel fully decentralized Reinforcement Learning (RL)-based solution that
achieves fairness without coordination. The proposed Fair Share RL (FSRL)
solution combines: (i) state augmentation with a semi-adaptive time reference;
(ii) an architecture that leverages risk control and time difference
likelihood; and (iii) a fairness-driven reward structure. We evaluate FSRL in
more than 50 network settings with different number of agents, different
amounts of available spectrum, in the presence of jammers, and in an ad-hoc
setting. Simulation results suggest that, when we compare FSRL with a common
baseline RL algorithm from the literature, FSRL can be up to 89.0% fairer (as
measured by Jain's fairness index) in stringent settings with several sources
and a single frequency band, and 48.1% fairer on average.


## AI 摘要

该研究提出了一种基于强化学习(RL)的完全去中心化无线网络资源分配方案FSRL，用于多个源-目的地对共享有限正交频带的场景。FSRL通过(i)半自适应时间参考的状态增强、(ii)结合风险控制和时差似然的架构、(iii)公平驱动的奖励结构，使各源节点在无需协调的情况下自主优化频带选择策略，仅通过观察自身传输结果(成功/冲突)学习，无需网络规模或其他节点策略的先验知识。在50多种网络设置下的仿真表明，相比基线RL算法，FSRL在严苛条件下公平性(Jain指数)最高提升89.0%，平均提升48.1%，同时保持高吞吐量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T01:29:28Z
- **目录日期**: 2025-04-02
