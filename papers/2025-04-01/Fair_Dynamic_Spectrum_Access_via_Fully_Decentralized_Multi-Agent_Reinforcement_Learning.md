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

该研究提出了一种完全去中心化的强化学习(FSRL)方案，用于解决多源-目的对共享有限正交频带的无线网络公平性问题。FSRL无需协调或信息共享，仅通过观察自身传输结果(成功/冲突)来学习频带选择策略。其创新点包括：(1)结合半自适应时间参考的状态增强；(2)利用风险控制和时差似然的架构；(3)公平导向的奖励机制。在50多种网络场景测试中(包括不同节点数、频谱资源、干扰和自组织环境)，FSRL相比基线算法在严苛条件下(多源单频带)公平性提升达89%(Jain指数)，平均提升48.1%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T05:02:01Z
- **目录日期**: 2025-04-01
