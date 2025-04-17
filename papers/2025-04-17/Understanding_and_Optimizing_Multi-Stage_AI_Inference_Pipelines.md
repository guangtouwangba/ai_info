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

随着大语言模型(LLMs)的快速发展，传统推理流程已无法满足多阶段任务需求。本文提出HERMES异构多阶段LLM推理模拟器，能精确建模检索增强生成(RAG)、KV缓存检索、动态路由等多阶段工作流。相比现有方案，HERMES支持异构硬件(CPU/GPU/ASIC)协同、多模型并发、先进批处理策略及多级内存层次，通过真实硬件轨迹与分析模型相结合，量化内存带宽争用、集群通信延迟等关键指标。案例研究表明，HERMES能有效评估推理阶段延迟、混合流水线批处理策略及远程KV缓存检索的架构影响，为下一代AI负载的软硬件协同设计提供优化指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-17T15:02:48Z
- **目录日期**: 2025-04-17
