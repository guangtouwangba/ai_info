# Realistic Urban Traffic Generator using Decentralized Federated Learning for the SUMO simulator

**URL**: http://arxiv.org/abs/2506.07980v1

## 原始摘要

Realistic urban traffic simulation is essential for sustainable urban
planning and the development of intelligent transportation systems. However,
generating high-fidelity, time-varying traffic profiles that accurately reflect
real-world conditions, especially in large-scale scenarios, remains a major
challenge. Existing methods often suffer from limitations in accuracy,
scalability, or raise privacy concerns due to centralized data processing. This
work introduces DesRUTGe (Decentralized Realistic Urban Traffic Generator), a
novel framework that integrates Deep Reinforcement Learning (DRL) agents with
the SUMO simulator to generate realistic 24-hour traffic patterns. A key
innovation of DesRUTGe is its use of Decentralized Federated Learning (DFL),
wherein each traffic detector and its corresponding urban zone function as an
independent learning node. These nodes train local DRL models using minimal
historical data and collaboratively refine their performance by exchanging
model parameters with selected peers (e.g., geographically adjacent zones),
without requiring a central coordinator. Evaluated using real-world data from
the city of Barcelona, DesRUTGe outperforms standard SUMO-based tools such as
RouteSampler, as well as other centralized learning approaches, by delivering
more accurate and privacy-preserving traffic pattern generation.


## AI 摘要

本文提出了一种名为DesRUTGe的新型城市交通模拟框架，通过结合深度强化学习(DRL)与SUMO模拟器，生成高保真的24小时交通模式。该框架采用去中心化联邦学习(DFL)，将每个交通探测器和对应城区作为独立学习节点，仅需少量历史数据训练本地DRL模型，并通过与相邻区域交换模型参数(而非原始数据)进行协作优化，无需中央协调器。基于巴塞罗那真实数据的测试表明，DesRUTGe在准确性和隐私保护方面优于传统SUMO工具和集中式学习方法，为大规模城市交通模拟提供了创新解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T18:01:42Z
- **目录日期**: 2025-06-10
