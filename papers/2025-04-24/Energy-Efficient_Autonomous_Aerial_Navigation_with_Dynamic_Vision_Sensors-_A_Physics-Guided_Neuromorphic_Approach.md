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

本文提出了一种基于神经形态视觉的无人机自主导航框架，通过事件相机（DVS）和脉冲神经网络（SNN）无监督检测动态障碍物，并结合轻量级物理引导神经网络（PgNN）预测最优飞行路径。该系统融合事件视觉与物理建模，在Gazebo仿真器中实现ROS集成的视觉-规划神经符号框架，具有高动态范围、抗运动模糊和低延迟决策优势。研究展示了事件视觉与物理引导规划相结合在节能自主导航（特别是动态门穿越避障场景）中的应用潜力，兼顾能效优化与符号可解释性。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T09:02:16Z
- **目录日期**: 2025-04-24
