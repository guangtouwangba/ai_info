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

本文提出了一种基于强化学习（RL）的自主水面艇（ASV）浅水导航框架。针对传统方法在动态干扰和深度限制下的不足，该方案仅需单波束测深仪（SBES）的稀疏深度数据，通过高斯过程（GP）回归逐步构建海底地形图，增强环境感知能力。结合RL算法，ASV能在有限传感器信息下安全高效地避开危险区域并抵达目标。实验验证了该方法在仿真与现实场景间的有效迁移性，显著提升了ASV在复杂浅水环境中的导航性能与安全性。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T02:31:40Z
- **目录日期**: 2025-06-04
