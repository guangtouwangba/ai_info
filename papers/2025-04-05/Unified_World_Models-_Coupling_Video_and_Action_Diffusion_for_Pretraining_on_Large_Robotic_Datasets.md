# Unified World Models: Coupling Video and Action Diffusion for Pretraining on Large Robotic Datasets

**URL**: http://arxiv.org/abs/2504.02792v1

## 原始摘要

Imitation learning has emerged as a promising approach towards building
generalist robots. However, scaling imitation learning for large robot
foundation models remains challenging due to its reliance on high-quality
expert demonstrations. Meanwhile, large amounts of video data depicting a wide
range of environments and diverse behaviors are readily available. This data
provides a rich source of information about real-world dynamics and
agent-environment interactions. Leveraging this data directly for imitation
learning, however, has proven difficult due to the lack of action annotation
required for most contemporary methods. In this work, we present Unified World
Models (UWM), a framework that allows for leveraging both video and action data
for policy learning. Specifically, a UWM integrates an action diffusion process
and a video diffusion process within a unified transformer architecture, where
independent diffusion timesteps govern each modality. We show that by simply
controlling each diffusion timestep, UWM can flexibly represent a policy, a
forward dynamics, an inverse dynamics, and a video generator. Through simulated
and real-world experiments, we show that: (1) UWM enables effective pretraining
on large-scale multitask robot datasets with both dynamics and action
predictions, resulting in more generalizable and robust policies than imitation
learning, (2) UWM naturally facilitates learning from action-free video data
through independent control of modality-specific diffusion timesteps, further
improving the performance of finetuned policies. Our results suggest that UWM
offers a promising step toward harnessing large, heterogeneous datasets for
scalable robot learning, and provides a simple unification between the often
disparate paradigms of imitation learning and world modeling. Videos and code
are available at https://weirdlabuw.github.io/uwm/.


## AI 摘要

本文提出了一种名为“统一世界模型（UWM）”的框架，通过整合动作扩散和视频扩散过程，利用大规模多模态数据（含动作标注和无动作视频）进行机器人策略学习。UWM基于统一Transformer架构，通过独立控制各模态的扩散时间步，灵活实现策略、正向/逆向动力学及视频生成功能。实验表明，UWM在仿真和真实场景中均优于传统模仿学习，能通过预训练提升策略的泛化性和鲁棒性，并有效利用无动作视频数据进一步优化策略。该框架为整合模仿学习与世界建模提供了新思路，支持异构数据驱动的可扩展机器人学习。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-05T13:05:38Z
- **目录日期**: 2025-04-05
