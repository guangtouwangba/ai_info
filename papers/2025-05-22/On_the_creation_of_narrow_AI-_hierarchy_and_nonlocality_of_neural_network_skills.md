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

该研究探讨了构建高效、安全的专用窄AI系统的两大挑战。首先，实验发现某些窄技能需要在大范围数据上训练才能掌握，因为技能间存在层级依赖关系，广泛数据分布能形成有效学习路径。其次，研究发现模型技能往往无法完全定位到特定可剪枝组件，但基于剪枝的方法仍优于蒸馏法。研究者提出使用正则化目标来对齐目标技能与可剪枝组件，同时消除不必要技能。这些发现为开发专用AI系统提供了重要见解，在效率和安全性方面具有潜在价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T09:00:59Z
- **目录日期**: 2025-05-22
