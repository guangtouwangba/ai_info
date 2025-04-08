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

Dion是一种高效的分布式AI训练优化器，通过正交化更新和本地动量缓冲技术大幅降低通信开销，同时保持同步训练的语义。与传统方法不同，Dion无需交换完整梯度矩阵，并采用高效的分片策略避免训练中重建大矩阵，显著减少I/O成本。该方法兼容标准分布式训练框架（如DDP、FSDP），在保持训练效果的同时提升了多加速器环境下的计算效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T17:01:52Z
- **目录日期**: 2025-04-08
