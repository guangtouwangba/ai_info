# Curvature Tuning: Provable Training-free Model Steering From a Single Parameter

**URL**: http://arxiv.org/abs/2502.07783v3

## 原始摘要

The scaling of model and data sizes has reshaped the AI landscape,
establishing finetuning pretrained models as the standard paradigm for solving
downstream tasks. However, dominant finetuning methods typically rely on weight
adaptation, often lack interpretability, and depend on heuristically chosen
hyperparameters. In this paper, we take a different perspective and shift the
focus from weights to activation functions, viewing them through the lens of
spline operators. We propose Curvature Tuning (CT), an interpretable and
principled steering method that modulates a model's decision boundary by
injecting a single hyperparameter into its activation functions. We show that
CT provably adjusts model decision boundary curvature and, more fundamentally,
projects a model onto a space of smooth functions-thereby complementing current
finetuning methods, whose effect lies primarily in feature adaptation. Making
this hyperparameter trainable gives rise to a novel and highly
parameter-efficient finetuning method. Empirically, CT improves both
generalization and robustness. For example, it boosts downstream accuracy of
ResNet-50/152 by 7.14%/8.46% over linear probing and 4.64%/1.70% over LoRA
across 12 datasets, and improves robust accuracy on the $\ell_\infty$ benchmark
from RobustBench by 1032.64%/1494.46%. Our code is available at
https://github.com/Leon-Leyang/curvature-tuning.


## AI 摘要

本文提出了一种新颖的微调方法——曲率调谐(CT)，通过调整激活函数的曲率而非权重来优化预训练模型。与传统权重微调方法相比，CT具有更好的可解释性，仅需引入单个超参数即可调节模型决策边界，并能将模型投影到平滑函数空间。实验表明，CT显著提升了模型性能：在12个数据集上，ResNet-50/152的下游准确率分别比线性探测提高了7.14%/8.46%，比LoRA方法提高了4.64%/1.70%；在RobustBench的鲁棒性测试中，准确率提升了1032.64%/1494.46%。该方法参数效率高，兼具泛化性和鲁棒性优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T09:02:42Z
- **目录日期**: 2025-06-11
