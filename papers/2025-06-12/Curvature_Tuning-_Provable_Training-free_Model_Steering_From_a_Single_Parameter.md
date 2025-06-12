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

这篇论文提出了一种新的微调方法——曲率调谐（Curvature Tuning, CT），通过调整激活函数的曲率来优化预训练模型。与传统基于权重调整的方法不同，CT将超参数注入激活函数，可证明地改变模型决策边界，提升模型平滑性。实验表明，CT显著提高了模型泛化能力和鲁棒性：在12个数据集上，ResNet-50/152的下游准确率分别比线性探测提高7.14%/8.46%，比LoRA提高4.64%/1.70%；在RobustBench的鲁棒性基准测试中，准确率提升高达1032.64%/1494.46%。该方法具有可解释性且参数高效。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T00:02:29Z
- **目录日期**: 2025-06-12
