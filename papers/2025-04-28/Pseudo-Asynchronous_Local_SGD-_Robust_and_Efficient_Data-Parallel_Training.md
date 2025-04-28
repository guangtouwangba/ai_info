# Pseudo-Asynchronous Local SGD: Robust and Efficient Data-Parallel Training

**URL**: http://arxiv.org/abs/2504.18454v1

## 原始摘要

Following AI scaling trends, frontier models continue to grow in size and
continue to be trained on larger datasets. Training these models requires huge
investments in exascale computational resources, which has in turn driven
development of distributed deep learning methods. Data parallelism is an
essential approach to speed up training, but it requires frequent global
communication between workers, which can bottleneck training at the largest
scales. In this work, we propose a method called Pseudo-Asynchronous Local SGD
(PALSGD) to improve the efficiency of data-parallel training. PALSGD is an
extension of Local SGD (Stich, 2018) and DiLoCo (Douillard et al., 2023),
designed to further reduce communication frequency by introducing a
pseudo-synchronization mechanism. PALSGD allows the use of longer
synchronization intervals compared to standard Local SGD. Despite the reduced
communication frequency, the pseudo-synchronization approach ensures that model
consistency is maintained, leading to performance results comparable to those
achieved with more frequent synchronization. Furthermore, we provide a
theoretical analysis of PALSGD, establishing its convergence and deriving its
convergence rate. This analysis offers insights into the algorithm's behavior
and performance guarantees. We evaluated PALSGD on image classification and
language modeling tasks. Our results show that PALSGD achieves better
performance in less time compared to existing methods like Distributed Data
Parallel (DDP), and DiLoCo. Notably, PALSGD trains 18.4% faster than DDP on
ImageNet-1K with ResNet-50, 24.4% faster than DDP on TinyStories with
GPT-Neo125M, and 21.1% faster than DDP on TinyStories with GPT-Neo-8M.


## AI 摘要

随着AI模型规模扩大，数据并行训练中的全局通信成为瓶颈。本研究提出Pseudo-Asynchronous Local SGD（PALSGD），通过伪同步机制减少通信频率，同时保持模型一致性。PALSGD扩展了Local SGD和DiLoCo，理论分析证明其收敛性。实验显示，在ImageNet-1K（ResNet-50）上比DDP快18.4%，在TinyStories（GPT-Neo125M和8M）上分别快24.4%和21.1%，性能相当但训练效率显著提升。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T23:01:40Z
- **目录日期**: 2025-04-28
