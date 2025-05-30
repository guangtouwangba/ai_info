# Adaptive Orchestration for Large-Scale Inference on Heterogeneous Accelerator Systems Balancing Cost, Performance, and Resilience

**URL**: http://arxiv.org/abs/2503.20074v2

## 原始摘要

The surge in generative AI workloads has created a need for scalable
inference systems that can flexibly harness both GPUs and specialized
accelerators while containing operational costs. This paper proposes a
hardware-agnostic control loop that adaptively allocates requests across
heterogeneous accelerators based on real-time cost and capacity signals. The
approach sustains low latency and high throughput by dynamically shifting
between cost-optimized and capacity-optimized modes, ensuring the most
efficient use of expensive compute resources under fluctuating availability.
Evaluated using the Stable Diffusion model, the framework consistently meets
latency targets, automatically redirects traffic during capacity shortfalls,
and capitalizes on lower-cost accelerators when possible. These results
highlight how a feedback-driven deployment strategy, spanning the entire
software and hardware stack, can help organizations efficiently scale
generative AI workloads while maintaining resilience in the face of limited
accelerator capacity.


## AI 摘要

本文提出了一种硬件无关的控制循环方法，用于在异构加速器上动态分配生成式AI推理请求。该系统通过实时监测成本和容量信号，在成本优化和容量优化模式间灵活切换，确保在资源波动时高效利用昂贵的计算资源。基于Stable Diffusion模型的评估显示，该框架能稳定满足延迟要求，在容量不足时自动转移流量，并优先使用低成本加速器。这种覆盖软硬件全栈的反馈驱动部署策略，可帮助组织在有限加速器资源下高效扩展生成式AI工作负载，同时保持系统韧性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T23:01:39Z
- **目录日期**: 2025-03-29
