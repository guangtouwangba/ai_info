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

本文提出了一种基于强化学习（RL）的自主水面艇（ASV）浅水导航框架，通过单波束测深仪（SBES）的稀疏深度测量实现安全高效航行。该框架结合高斯过程（GP）回归，逐步构建海底地形图以增强环境感知能力，从而优化决策。实验验证了该方法在动态干扰和深度限制下的有效性，并实现了从仿真到实际环境的有效迁移，确保在复杂浅水环境中保持安全性。研究结果表明，该RL框架显著提升了ASV在有限传感器信息条件下的导航性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T09:01:52Z
- **目录日期**: 2025-06-03
