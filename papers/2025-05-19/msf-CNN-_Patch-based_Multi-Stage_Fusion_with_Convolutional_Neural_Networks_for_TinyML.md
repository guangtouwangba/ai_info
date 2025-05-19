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

本文介绍了msf-CNN，一种针对微控制器（MCU）优化的新型卷积神经网络融合技术。该方法通过将融合解空间表示为有向无环图，高效寻找最优融合方案，相比现有技术（如MCUNetV2和StreamNet）能发现更多解决方案。实验表明，msf-CNN在ARM Cortex-M、RISC-V和ESP32等MCU上运行时，可减少50%的RAM使用量，同时保持实时推理性能。这项技术为系统设计者提供了更大的灵活性，帮助在MCU极小的内存预算（如128kB RAM）下实现高效的AI模型部署。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T06:00:57Z
- **目录日期**: 2025-05-19
