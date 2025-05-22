# On the creation of narrow AI: hierarchy and nonlocality of neural network skills

**URL**: http://arxiv.org/abs/2505.15811v1

## 原始摘要

We study the problem of creating strong, yet narrow, AI systems. While recent
AI progress has been driven by the training of large general-purpose foundation
models, the creation of smaller models specialized for narrow domains could be
valuable for both efficiency and safety. In this work, we explore two
challenges involved in creating such systems, having to do with basic
properties of how neural networks learn and structure their representations.
The first challenge regards when it is possible to train narrow models from
scratch. Through experiments on a synthetic task, we find that it is sometimes
necessary to train networks on a wide distribution of data to learn certain
narrow skills within that distribution. This effect arises when skills depend
on each other hierarchically, and training on a broad distribution introduces a
curriculum which substantially accelerates learning. The second challenge
regards how to transfer particular skills from large general models into small
specialized models. We find that model skills are often not perfectly localized
to a particular set of prunable components. However, we find that methods based
on pruning can still outperform distillation. We investigate the use of a
regularization objective to align desired skills with prunable components while
unlearning unnecessary skills.


## AI 摘要

这项研究探讨了开发高效、安全的专用窄AI系统的两大挑战。首先，实验发现某些窄技能需要基于广泛数据训练才能掌握，尤其是当技能间存在层级依赖时，宽分布数据能提供加速学习的内在课程。其次，研究考察了如何从通用大模型迁移特定技能到专用小模型，发现技能通常无法完美定位到特定可剪枝组件，但剪枝方法仍优于蒸馏法。研究者还探索了通过正则化目标使所需技能与可剪枝组件对齐，同时遗忘不必要技能的方法。这些发现为开发更高效的专用AI系统提供了重要见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T21:01:03Z
- **目录日期**: 2025-05-22
