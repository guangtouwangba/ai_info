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

本文提出了一种名为msf-CNN的新型技术，通过将融合解空间表示为有向无环图，高效寻找卷积神经网络(CNN)的最佳融合设置。该方法针对微控制器(MCU)的内存限制(如128kB RAM)和实时性要求，采用基于块的融合优化数据流。相比现有方法(MCUNetV2和StreamNet)，msf-CNN能发现更多解决方案，并在多种MCU平台(ARM Cortex-M、RISC-V、ESP32)上实现运行，内存使用减少50%。该技术为系统设计者提供了更大的灵活性，在保持低延迟的同时显著提升了内存效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-19T23:01:03Z
- **目录日期**: 2025-05-19
