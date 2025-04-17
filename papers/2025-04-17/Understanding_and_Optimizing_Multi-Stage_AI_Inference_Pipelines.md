# Understanding and Optimizing Multi-Stage AI Inference Pipelines

**URL**: http://arxiv.org/abs/2504.09775v2

## 原始摘要

The rapid evolution of Large Language Models (LLMs) has driven the need for
increasingly sophisticated inference pipelines and hardware platforms. Modern
LLM serving extends beyond traditional prefill-decode workflows, incorporating
multi-stage processes such as Retrieval Augmented Generation (RAG), key-value
(KV) cache retrieval, dynamic model routing, and multi step reasoning. These
stages exhibit diverse computational demands, requiring distributed systems
that integrate GPUs, ASICs, CPUs, and memory-centric architectures. However,
existing simulators lack the fidelity to model these heterogeneous,
multi-engine workflows, limiting their ability to inform architectural
decisions.
  To address this gap, we introduce HERMES, a Heterogeneous Multi-stage LLM
inference Execution Simulator. HERMES models diverse request stages; including
RAG, KV retrieval, reasoning, prefill, and decode across complex hardware
hierarchies. HERMES supports heterogeneous clients executing multiple models
concurrently unlike prior frameworks while incorporating advanced batching
strategies and multi-level memory hierarchies. By integrating real hardware
traces with analytical modeling, HERMES captures critical trade-offs such as
memory bandwidth contention, inter-cluster communication latency, and batching
efficiency in hybrid CPU-accelerator deployments. Through case studies, we
explore the impact of reasoning stages on end-to-end latency, optimal batching
strategies for hybrid pipelines, and the architectural implications of remote
KV cache retrieval. HERMES empowers system designers to navigate the evolving
landscape of LLM inference, providing actionable insights into optimizing
hardware-software co-design for next-generation AI workloads.


## AI 摘要

随着大语言模型(LLM)的快速发展，传统推理流程已无法满足RAG、KV缓存检索等新型多阶段工作负载的需求。现有模拟器难以准确建模这些异构计算场景。为此，研究者提出HERMES异构多阶段LLM推理模拟器，支持RAG、KV检索、推理、预填充和解码等全流程模拟，能精确刻画CPU-加速器混合架构中的内存带宽争用、批处理效率等关键指标。通过案例研究，HERMES揭示了推理阶段对时延的影响、最优批处理策略等洞见，为下一代AI工作负载的软硬件协同设计提供了重要工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T19:01:52Z
- **目录日期**: 2025-04-17
