# Dion: A Communication-Efficient Optimizer for Large Models

**URL**: http://arxiv.org/abs/2504.05295v1

## 原始摘要

Training large AI models efficiently requires distributing computation across
multiple accelerators, but this often incurs significant communication overhead
-- especially during gradient synchronization. We introduce Dion, a
communication-efficient optimizer that retains the synchronous semantics of
standard distributed training (e.g., DDP, FSDP) while substantially reducing
I/O costs. Unlike conventional optimizers that synchronize full gradient
matrices, Dion leverages orthonormalized updates with device-local momentum
buffers, eliminating the need for full gradient exchange. It further supports
an efficient sharding strategy that avoids reconstructing large matrices during
training.


## AI 摘要

Dion是一种高效的分布式AI训练优化器，通过正交化更新和本地动量缓冲技术，大幅减少了梯度同步时的通信开销。它在保持标准分布式训练（如DDP、FSDP）同步语义的同时，避免了全梯度矩阵交换。Dion还支持高效的分片策略，防止训练过程中重构大矩阵，显著降低了I/O成本。相比传统优化器，Dion在维持模型性能的前提下，显著提升了大规模AI模型在多加速器上的训练效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T02:27:59Z
- **目录日期**: 2025-04-09
