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

该研究探讨了开发高效、安全的专用窄AI系统的两个关键挑战。首先，实验发现某些窄技能需要基于广泛数据训练才能掌握，尤其是当技能存在层级依赖时，宽分布数据能形成内在学习梯度从而加速训练。其次，研究发现从通用大模型向专用小模型迁移技能时，技能往往无法完全定位到特定可剪枝组件中，但基于剪枝的方法仍优于蒸馏法。研究还探索了通过正则化目标来对齐目标技能与可剪枝组件，同时消除不必要技能的方法。这些发现为开发专业化AI系统提供了重要见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T08:01:01Z
- **目录日期**: 2025-05-22
