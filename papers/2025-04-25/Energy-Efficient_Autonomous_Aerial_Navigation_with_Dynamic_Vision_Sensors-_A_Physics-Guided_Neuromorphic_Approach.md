# Energy-Efficient Autonomous Aerial Navigation with Dynamic Vision Sensors: A Physics-Guided Neuromorphic Approach

**URL**: http://arxiv.org/abs/2502.05938v2

## 原始摘要

Vision-based object tracking is a critical component for achieving autonomous
aerial navigation, particularly for obstacle avoidance. Neuromorphic Dynamic
Vision Sensors (DVS) or event cameras, inspired by biological vision, offer a
promising alternative to conventional frame-based cameras. These cameras can
detect changes in intensity asynchronously, even in challenging lighting
conditions, with a high dynamic range and resistance to motion blur. Spiking
neural networks (SNNs) are increasingly used to process these event-based
signals efficiently and asynchronously. Meanwhile, physics-based artificial
intelligence (AI) provides a means to incorporate system-level knowledge into
neural networks via physical modeling. This enhances robustness, energy
efficiency, and provides symbolic explainability. In this work, we present a
neuromorphic navigation framework for autonomous drone navigation. The focus is
on detecting and navigating through moving gates while avoiding collisions. We
use event cameras for detecting moving objects through a shallow SNN
architecture in an unsupervised manner. This is combined with a lightweight
energy-aware physics-guided neural network (PgNN) trained with depth inputs to
predict optimal flight times, generating near-minimum energy paths. The system
is implemented in the Gazebo simulator and integrates a sensor-fused
vision-to-planning neuro-symbolic framework built with the Robot Operating
System (ROS) middleware. This work highlights the future potential of
integrating event-based vision with physics-guided planning for
energy-efficient autonomous navigation, particularly for low-latency
decision-making.


## AI 摘要

该研究提出了一种基于神经形态动态视觉传感器（DVS）和脉冲神经网络（SNN）的无人机自主导航框架。通过事件相机在无监督模式下检测移动障碍物，并结合轻量级物理引导神经网络（PgNN）预测最优飞行路径以实现近最小能耗。系统在Gazebo模拟器中实现，采用ROS中间件构建了从视觉到规划的神经符号框架。该方案融合了基于事件的视觉感知与物理建模，在复杂光照条件下具有高动态范围和抗运动模糊优势，为低延迟决策的节能自主导航提供了新思路，特别适用于动态障碍物规避场景。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T02:40:12Z
- **目录日期**: 2025-04-25
