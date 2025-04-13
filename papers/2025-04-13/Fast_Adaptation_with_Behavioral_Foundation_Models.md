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

无监督零样本强化学习（RL）通过预训练行为基础模型（BFM），使智能体无需额外学习即可解决下游任务。然而，由于训练误差和推理过程，零样本策略往往不够理想。研究发现，BFM学习的技能库中存在比推理结果更优的策略，适合快速适应。为此，研究者提出了基于行动者-评论家和纯行动者的快速适应策略，在低维任务嵌入空间中优化零样本策略性能，避免了预训练模型的初始性能下降问题。实验表明，该方法在导航和运动任务中仅需几十次交互即可提升10-40%的性能，优于现有基线。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T09:01:28Z
- **目录日期**: 2025-04-13
