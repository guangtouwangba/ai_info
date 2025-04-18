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

本文提出了一种基于强化学习(RL)的完全去中心化无线网络公平共享方案(FSRL)。该方案使多个源-目的地对能在不共享信息的情况下，仅通过观察自身传输结果(成功/冲突)，自适应地选择频段。FSRL结合了：(1)半自适应时间参考的状态增强；(2)利用风险控制和时差似然的架构；(3)公平驱动的奖励机制。在50多种网络设置下的测试表明，与基线RL算法相比，FSRL在严格条件下(多源单频段)公平性提升达89.0%，平均提升48.1%(以Jain公平指数衡量)，有效实现了网络范围内的公平性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T23:02:02Z
- **目录日期**: 2025-04-01
