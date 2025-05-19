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

本文介绍了msf-CNN技术，这是一种针对微控制器(MCU)优化的卷积神经网络(CNN)融合方法。通过将融合解空间表示为有向无环图，msf-CNN能更高效地找到最优融合方案，显著降低内存需求。实验表明，相比现有技术(MCUNetV2和StreamNet)，msf-CNN可减少50%的RAM使用量，同时保持实时推理性能。该方法已实现在多种MCU架构(ARM Cortex-M、RISC-V、ESP32)上运行，为系统设计者提供了更大的灵活性，使AI模型能在资源受限的MCU(如仅128kB RAM)上高效部署。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T07:00:59Z
- **目录日期**: 2025-05-19
