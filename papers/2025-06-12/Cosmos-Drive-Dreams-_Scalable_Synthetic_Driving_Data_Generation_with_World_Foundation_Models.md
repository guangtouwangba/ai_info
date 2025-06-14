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

NVIDIA推出的Cosmos-Drive-Dreams是一个合成数据生成（SDG）管道，旨在为自动驾驶车辆（AV）等安全关键AI系统生成具有挑战性的场景。该管道基于Cosmos-Drive模型套件，能够生成可控、高保真、多视角且时空一致的驾驶视频。通过生成高保真和多样化的驾驶场景数据，该工具解决了现实数据收集成本高且难以捕捉边缘案例的问题。实验表明，生成的数据有助于缓解长尾分布问题，并提升3D车道检测、3D物体检测和驾驶策略学习等下游任务的泛化能力。相关工具、数据集和模型权重已在NVIDIA Cosmos平台开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T01:29:16Z
- **目录日期**: 2025-06-12
