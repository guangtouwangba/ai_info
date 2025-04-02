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

本文提出了一种去中心化无线网络中基于强化学习(RL)的公平共享方案FSRL。该方案使多个源节点能在不共享信息、不知晓网络规模的情况下，通过自主观察传输结果(成功/碰撞)来调整频段选择策略。FSRL结合了状态增强、半自适应时间参考、风险控制架构和公平驱动的奖励机制，在50多种网络设置中表现出色。与基线RL算法相比，FSRL在严苛条件下(多源节点单频段)公平性提升达89%(按Jain公平指数)，平均提升48.1%，实现了网络范围内的公平性同时最大化各节点吞吐量，且无需协调。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T02:28:59Z
- **目录日期**: 2025-04-02
