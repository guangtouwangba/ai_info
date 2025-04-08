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

Dion是一种高效的分布式训练优化器，通过正交化更新和本地动量缓冲大幅减少I/O开销，同时保持标准分布式训练（如DDP、FSDP）的同步语义。与传统优化器不同，Dion无需交换完整梯度矩阵，并采用高效的分片策略避免训练期间重建大矩阵，显著降低了通信成本。该方法在保持训练效果的同时优化了多加速器间的计算分配效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T06:01:26Z
- **目录日期**: 2025-04-08
