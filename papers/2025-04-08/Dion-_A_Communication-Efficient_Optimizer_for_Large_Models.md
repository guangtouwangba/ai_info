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

Dion是一种高效的分布式AI训练优化器，通过正交化更新和本地动量缓冲减少通信开销，同时保持同步训练的语义。它避免了传统优化器中全梯度矩阵的同步，显著降低了I/O成本。此外，Dion支持高效的分片策略，避免训练过程中重建大矩阵，从而提升大规模模型训练的效率和可扩展性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T21:01:58Z
- **目录日期**: 2025-04-08
