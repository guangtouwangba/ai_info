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

这篇论文提出了一种硬件无关的控制循环方法，用于动态分配生成式AI推理任务到异构加速器（如GPU和专用加速器）上。该系统通过实时监测成本和容量信号，在成本优化和容量优化模式间灵活切换，从而在资源波动下维持低延迟和高吞吐量。基于Stable Diffusion模型的评估显示，该框架能稳定满足延迟要求，在资源不足时自动转移流量，并优先使用低成本加速器。这种反馈驱动的部署策略有助于组织高效扩展生成式AI工作负载，在加速器资源有限的情况下保持系统韧性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-28T22:01:27Z
- **目录日期**: 2025-03-28
