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

Dion是一种高效的分布式训练优化器，通过正交化更新和本地动量缓冲减少通信开销，同时保持同步训练语义（如DDP/FSDP）。与传统方法需同步完整梯度矩阵不同，Dion利用设备本地计算避免全量梯度交换，并采用分片策略防止大矩阵重构。实验显示，在保持模型性能的同时，Dion显著降低了I/O成本，通信量减少达85%，吞吐量提升1.3-2.4倍，尤其适合大规模AI模型训练。该方法兼容主流框架（PyTorch），无需修改模型架构即可部署。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-08T20:01:51Z
- **目录日期**: 2025-04-08
