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

本文介绍了msf-CNN，一种针对微控制器（MCU）的新型卷积神经网络优化技术。该方法通过基于有向无环图的融合解决方案空间搜索，显著提升了内存效率，在ARM Cortex-M、RISC-V和ESP32等MCU上运行时，比现有技术（MCUNetV2和StreamNet）减少50%的RAM使用量。msf-CNN采用基于补丁的融合策略优化神经网络层间数据流，在保持实时推理性能的同时，为系统设计者提供了更大的灵活性。该技术特别适合内存资源受限（如128kB RAM）的微型设备，扩展了AI在边缘计算中的应用可能性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T11:00:58Z
- **目录日期**: 2025-05-19
