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

Dion是一种高效的分布式AI训练优化器，通过正交化更新和本地动量缓冲减少通信开销，同时保持同步训练语义（如DDP/FSDP）。与传统方法需同步完整梯度矩阵不同，Dion仅交换低维正交基，显著降低I/O成本。它还支持分片策略，避免训练中重建大矩阵。实验显示，Dion在保持模型性能的同时，通信量比FSDP减少高达4.6倍，训练速度提升1.25-1.53倍，内存占用降低1.6-3.3倍，适用于不同规模模型（1B-70B参数）。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T00:01:29Z
- **目录日期**: 2025-04-09
