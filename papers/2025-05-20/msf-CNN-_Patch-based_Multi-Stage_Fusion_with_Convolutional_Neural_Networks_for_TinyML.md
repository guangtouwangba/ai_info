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

本文介绍了msf-CNN，一种针对微控制器（MCU）优化的新型卷积神经网络融合技术。该方法通过将融合解空间建模为有向无环图，高效搜索最优融合配置，显著扩展了可行解范围。相比现有方案（如MCUNetV2和StreamNet），msf-CNN在保持实时推理性能的同时，可减少50%的RAM占用。研究团队已实现支持多种MCU架构（ARM Cortex-M、RISC-V、ESP32）的开源方案，为系统设计者提供了更灵活的内存-延迟权衡空间。该技术特别适用于内存严格受限（如128kB RAM）的嵌入式场景，通过优化跨层数据流提升模型部署效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T01:28:44Z
- **目录日期**: 2025-05-20
