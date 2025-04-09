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

Dion是一种高效的分布式训练优化器，通过正交化更新和本地动量缓冲减少通信开销，同时保持同步训练语义（如DDP/FSDP）。它避免了传统优化器同步完整梯度矩阵的需求，并采用分片策略防止训练中重构大矩阵，显著降低I/O成本。该方法在保持训练效果的同时，提升了多加速器环境下大规模AI模型的训练效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T03:13:45Z
- **目录日期**: 2025-04-09
