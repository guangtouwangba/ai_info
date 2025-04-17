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

随着大语言模型(LLMs)的快速发展，其推理流程日益复杂，涉及检索增强生成(RAG)、KV缓存检索、动态模型路由等多阶段任务，对异构硬件系统提出了更高要求。现有模拟器难以准确建模这种多引擎工作流。为此，研究者提出了HERMES异构多阶段LLM推理模拟器，它能精确模拟RAG、KV检索、推理等各阶段在不同硬件架构上的执行情况，支持多模型并发、高级批处理策略和多级内存层次。通过结合真实硬件追踪和分析建模，HERMES能评估内存带宽争用、集群通信延迟等关键权衡，为下一代AI工作负载的软硬件协同设计提供优化指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T17:01:59Z
- **目录日期**: 2025-04-17
