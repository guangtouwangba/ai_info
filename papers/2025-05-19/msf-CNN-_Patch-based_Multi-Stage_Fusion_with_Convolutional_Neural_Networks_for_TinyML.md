# msf-CNN: Patch-based Multi-Stage Fusion with Convolutional Neural Networks for TinyML

**URL**: http://arxiv.org/abs/2505.11483v1

## 原始摘要

AI spans from large language models to tiny models running on
microcontrollers (MCUs). Extremely memory-efficient model architectures are
decisive to fit within an MCU's tiny memory budget e.g., 128kB of RAM. However,
inference latency must remain small to fit real-time constraints. An approach
to tackle this is patch-based fusion, which aims to optimize data flows across
neural network layers. In this paper, we introduce msf-CNN, a novel technique
that efficiently finds optimal fusion settings for convolutional neural
networks (CNNs) by walking through the fusion solution space represented as a
directed acyclic graph. Compared to previous work on CNN fusion for MCUs,
msf-CNN identifies a wider set of solutions. We published an implementation of
msf-CNN running on various microcontrollers (ARM Cortex-M, RISC-V, ESP32). We
show that msf-CNN can achieve inference using 50% less RAM compared to the
prior art (MCUNetV2 and StreamNet). We thus demonstrate how msf-CNN offers
additional flexibility for system designers.


## AI 摘要

本文介绍了msf-CNN，一种针对微控制器（MCU）优化的卷积神经网络融合技术。通过将融合解空间建模为有向无环图，该方法能高效探索更多优化方案，显著降低内存占用。实验表明，相比MCUNetV2和StreamNet等现有方案，msf-CNN可减少50%的RAM使用量，同时满足实时性要求。该技术已实现跨平台支持（ARM Cortex-M/RISC-V/ESP32），为系统设计者提供了更灵活的内存-延迟权衡空间。这项突破使复杂AI模型能部署在仅128kB RAM的微型设备上，推动了边缘计算发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T15:00:55Z
- **目录日期**: 2025-05-19
