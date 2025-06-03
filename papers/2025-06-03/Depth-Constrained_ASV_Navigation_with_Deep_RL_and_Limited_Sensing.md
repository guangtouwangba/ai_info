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

本文提出了一种基于强化学习（RL）的自主水面艇（ASV）浅水导航框架，通过单波束测深仪（SBES）的稀疏深度测量实现安全航行。该方法结合高斯过程（GP）回归，逐步构建海底地形图以增强环境感知能力，从而优化决策。研究验证了该框架在仿真与真实场景间的有效迁移能力，确保训练策略在实际水域中的泛化性。实验结果表明，该方法在动态干扰和深度限制的挑战性浅水环境中，既能提升ASV导航性能，又能保障航行安全。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T21:02:15Z
- **目录日期**: 2025-06-03
