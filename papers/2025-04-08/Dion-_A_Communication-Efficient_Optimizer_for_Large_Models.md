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

Dion是一种高效通信优化器，用于分布式AI模型训练。它通过正交化更新和本地动量缓冲，避免了传统优化器同步完整梯度矩阵的开销，显著降低I/O成本。Dion保持了标准分布式训练（如DDP、FSDP）的同步语义，同时支持高效的参数分片策略，避免训练中重建大型矩阵。相比现有方法，Dion在减少通信开销的同时，维持了训练效果，为大规模模型训练提供了更高效的解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T07:01:03Z
- **目录日期**: 2025-04-08
