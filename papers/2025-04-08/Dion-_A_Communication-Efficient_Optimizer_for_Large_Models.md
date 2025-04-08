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

Dion是一种高效的分布式AI训练优化器，通过正交化更新和本地动量缓冲技术，显著减少梯度同步时的I/O开销。它保持了标准分布式训练(如DDP、FSDP)的同步语义，但无需交换完整梯度矩阵。Dion还支持高效的分片策略，避免训练过程中重建大型矩阵，从而降低通信成本。相比传统优化器同步完整梯度矩阵的方法，Dion在保持训练效果的同时大幅提升了多加速器环境下的训练效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T13:08:46Z
- **目录日期**: 2025-04-08
