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

本文介绍了msf-CNN，一种针对微控制器（MCU）优化的新型卷积神经网络融合技术。该方法通过将融合解空间表示为有向无环图，有效寻找最优融合配置，能在128kB RAM等严格内存限制下运行。相比现有方案（MCUNetV2和StreamNet），msf-CNN发现了更广泛的解决方案，内存占用减少50%，同时保持实时推理的低延迟。研究团队已发布支持多种MCU架构（ARM Cortex-M、RISC-V、ESP32）的实现方案，为系统设计者提供了更大的灵活性。该技术通过优化神经网络层间数据流，显著提升了MCU环境下的AI模型部署效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T14:00:59Z
- **目录日期**: 2025-05-19
