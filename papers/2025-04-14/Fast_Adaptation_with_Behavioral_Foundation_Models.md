# Fast Adaptation with Behavioral Foundation Models

**URL**: http://arxiv.org/abs/2504.07896v1

## 原始摘要

Unsupervised zero-shot reinforcement learning (RL) has emerged as a powerful
paradigm for pretraining behavioral foundation models (BFMs), enabling agents
to solve a wide range of downstream tasks specified via reward functions in a
zero-shot fashion, i.e., without additional test-time learning or planning.
This is achieved by learning self-supervised task embeddings alongside
corresponding near-optimal behaviors and incorporating an inference procedure
to directly retrieve the latent task embedding and associated policy for any
given reward function. Despite promising results, zero-shot policies are often
suboptimal due to errors induced by the unsupervised training process, the
embedding, and the inference procedure. In this paper, we focus on devising
fast adaptation strategies to improve the zero-shot performance of BFMs in a
few steps of online interaction with the environment while avoiding any
performance drop during the adaptation process. Notably, we demonstrate that
existing BFMs learn a set of skills containing more performant policies than
those identified by their inference procedure, making them well-suited for fast
adaptation. Motivated by this observation, we propose both actor-critic and
actor-only fast adaptation strategies that search in the low-dimensional
task-embedding space of the pre-trained BFM to rapidly improve the performance
of its zero-shot policies on any downstream task. Notably, our approach
mitigates the initial "unlearning" phase commonly observed when fine-tuning
pre-trained RL models. We evaluate our fast adaptation strategies on top of
four state-of-the-art zero-shot RL methods in multiple navigation and
locomotion domains. Our results show that they achieve 10-40% improvement over
their zero-shot performance in a few tens of episodes, outperforming existing
baselines.


## AI 摘要

这篇论文提出了一种快速适应策略，用于提升无监督零样本强化学习（RL）中行为基础模型（BFMs）的性能。研究发现，现有的BFMs已学习到比其推理过程识别到的更优策略，但零样本策略常因训练误差而表现欠佳。作者提出了基于行动者-评论家和纯行动者的低维任务嵌入空间搜索方法，能在少量环境交互步骤中快速提升下游任务表现，避免了预训练模型微调初期常见的性能下降问题。实验表明，该方法在导航和运动任务中仅需数十次迭代即可将零样本性能提升10-40%，优于现有基线。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T03:16:34Z
- **目录日期**: 2025-04-14
