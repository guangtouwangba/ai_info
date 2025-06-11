# Cosmos-Drive-Dreams: Scalable Synthetic Driving Data Generation with World Foundation Models

**URL**: http://arxiv.org/abs/2506.09042v1

## 原始摘要

Collecting and annotating real-world data for safety-critical physical AI
systems, such as Autonomous Vehicle (AV), is time-consuming and costly. It is
especially challenging to capture rare edge cases, which play a critical role
in training and testing of an AV system. To address this challenge, we
introduce the Cosmos-Drive-Dreams - a synthetic data generation (SDG) pipeline
that aims to generate challenging scenarios to facilitate downstream tasks such
as perception and driving policy training. Powering this pipeline is
Cosmos-Drive, a suite of models specialized from NVIDIA Cosmos world foundation
model for the driving domain and are capable of controllable, high-fidelity,
multi-view, and spatiotemporally consistent driving video generation. We
showcase the utility of these models by applying Cosmos-Drive-Dreams to scale
the quantity and diversity of driving datasets with high-fidelity and
challenging scenarios. Experimentally, we demonstrate that our generated data
helps in mitigating long-tail distribution problems and enhances generalization
in downstream tasks such as 3D lane detection, 3D object detection and driving
policy learning. We open source our pipeline toolkit, dataset and model weights
through the NVIDIA's Cosmos platform.
  Project page: https://research.nvidia.com/labs/toronto-ai/cosmos_drive_dreams


## AI 摘要

NVIDIA的研究团队开发了Cosmos-Drive-Dreams合成数据生成管道，旨在为自动驾驶系统（AV）创建高保真、多视角且时空一致的挑战性场景数据。该工具基于Cosmos-Drive模型套件，能够生成可控的驾驶视频，解决真实数据收集成本高且难以捕捉罕见边缘案例的问题。实验表明，生成的数据有助于缓解长尾分布问题，并提升3D车道检测、3D物体检测和驾驶策略学习等下游任务的泛化能力。团队已通过NVIDIA Cosmos平台开源了相关工具、数据集和模型权重。项目详情见研究页面。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T12:01:47Z
- **目录日期**: 2025-06-11
