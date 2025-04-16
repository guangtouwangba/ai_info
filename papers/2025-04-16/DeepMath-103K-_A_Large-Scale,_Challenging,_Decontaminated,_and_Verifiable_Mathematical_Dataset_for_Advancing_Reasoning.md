# DeepMath-103K: A Large-Scale, Challenging, Decontaminated, and Verifiable Mathematical Dataset for Advancing Reasoning

**URL**: http://arxiv.org/abs/2504.11456v1

## 原始摘要

The capacity for complex mathematical reasoning is a key benchmark for
artificial intelligence. While reinforcement learning (RL) applied to LLMs
shows promise, progress is significantly hindered by the lack of large-scale
training data that is sufficiently challenging, possesses verifiable answer
formats suitable for RL, and is free from contamination with evaluation
benchmarks. To address these limitations, we introduce DeepMath-103K, a new,
large-scale dataset comprising approximately 103K mathematical problems,
specifically designed to train advanced reasoning models via RL. DeepMath-103K
is curated through a rigorous pipeline involving source analysis, stringent
decontamination against numerous benchmarks, and filtering for high difficulty
(primarily Levels 5-9), significantly exceeding existing open resources in
challenge. Each problem includes a verifiable final answer, enabling rule-based
RL, and three distinct R1-generated solutions suitable for diverse training
paradigms like supervised fine-tuning or distillation. Spanning a wide range of
mathematical topics, DeepMath-103K promotes the development of generalizable
reasoning. We demonstrate that models trained on DeepMath-103K achieve
significant improvements on challenging mathematical benchmarks, validating its
effectiveness. We release DeepMath-103K publicly to facilitate community
progress in building more capable AI reasoning systems:
https://github.com/zwhe99/DeepMath.


## AI 摘要

DeepMath-103K是一个专为强化学习（RL）设计的大规模数学推理数据集，包含约10.3万道高难度数学题（难度等级5-9），远超现有开源资源。该数据集通过严格筛选、去污处理确保无评测基准污染，并提供可验证的最终答案和三种AI生成的解决方案，适用于监督微调或蒸馏等训练范式。DeepMath-103K覆盖广泛数学主题，能有效提升模型的复杂推理能力。实验证明，基于该数据集训练的模型在数学评测中表现显著提升。数据集已开源以推动AI推理系统发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T23:01:33Z
- **目录日期**: 2025-04-16
