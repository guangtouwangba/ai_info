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

这篇论文提出了一种硬件无关的控制循环系统，用于动态分配生成式AI推理任务到异构加速器（如GPU和专用芯片）上。该系统通过实时监控成本和容量信号，在成本优化和容量优化模式间灵活切换，确保在高负载下维持低延迟和高吞吐。基于Stable Diffusion模型的测试表明，该框架能自动应对容量短缺，优先使用低成本加速器，同时满足延迟要求。这种基于反馈的部署策略实现了计算资源的高效利用，帮助企业在加速器资源有限的情况下，经济高效地扩展生成式AI工作负载。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-30T22:01:42Z
- **目录日期**: 2025-03-30
