# Depth-Constrained ASV Navigation with Deep RL and Limited Sensing

**URL**: http://arxiv.org/abs/2504.18253v2

## 原始摘要

Autonomous Surface Vehicles (ASVs) play a crucial role in maritime
operations, yet their navigation in shallow-water environments remains
challenging due to dynamic disturbances and depth constraints. Traditional
navigation strategies struggle with limited sensor information, making safe and
efficient operation difficult. In this paper, we propose a reinforcement
learning (RL) framework for ASV navigation under depth constraints, where the
vehicle must reach a target while avoiding unsafe areas with only a single
depth measurement per timestep from a downward-facing Single Beam Echosounder
(SBES). To enhance environmental awareness, we integrate Gaussian Process (GP)
regression into the RL framework, enabling the agent to progressively estimate
a bathymetric depth map from sparse sonar readings. This approach improves
decision-making by providing a richer representation of the environment.
Furthermore, we demonstrate effective sim-to-real transfer, ensuring that
trained policies generalize well to real-world aquatic conditions. Experimental
results validate our method's capability to improve ASV navigation performance
while maintaining safety in challenging shallow-water environments.


## AI 摘要

本文提出了一种基于强化学习(RL)的自主水面艇(ASV)浅水导航框架。针对传统方法在传感器信息有限情况下的不足，该方案利用单波束测深仪(SBES)的稀疏深度测量数据，结合高斯过程(GP)回归逐步构建海底地形图，增强环境感知能力。研究实现了从仿真到实际环境的有效迁移，验证了该方法在动态干扰和深度约束的复杂浅水环境中，既能保证航行安全，又能提高导航性能。实验结果表明，该RL框架仅需单次深度测量即可支持ASV安全高效地抵达目标位置。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T10:01:43Z
- **目录日期**: 2025-06-03
