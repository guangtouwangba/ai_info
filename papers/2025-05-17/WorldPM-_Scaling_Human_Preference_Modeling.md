# WorldPM: Scaling Human Preference Modeling

**URL**: http://arxiv.org/abs/2505.10527v1

## 原始摘要

Motivated by scaling laws in language modeling that demonstrate how test loss
scales as a power law with model and dataset sizes, we find that similar laws
exist in preference modeling. We propose World Preference Modeling$ (WorldPM)
to emphasize this scaling potential, where World Preference embodies a unified
representation of human preferences. In this paper, we collect preference data
from public forums covering diverse user communities, and conduct extensive
training using 15M-scale data across models ranging from 1.5B to 72B
parameters. We observe distinct patterns across different evaluation metrics:
(1) Adversarial metrics (ability to identify deceptive features) consistently
scale up with increased training data and base model size; (2) Objective
metrics (objective knowledge with well-defined answers) show emergent behavior
in larger language models, highlighting WorldPM's scalability potential; (3)
Subjective metrics (subjective preferences from a limited number of humans or
AI) do not demonstrate scaling trends. Further experiments validate the
effectiveness of WorldPM as a foundation for preference fine-tuning. Through
evaluations on 7 benchmarks with 20 subtasks, we find that WorldPM broadly
improves the generalization performance across human preference datasets of
varying sizes (7K, 100K and 800K samples), with performance gains exceeding 5%
on many key subtasks. Integrating WorldPM into our internal RLHF pipeline, we
observe significant improvements on both in-house and public evaluation sets,
with notable gains of 4% to 8% in our in-house evaluations.


## AI 摘要

本文提出世界偏好建模(WorldPM)，通过分析语言模型中的幂律缩放现象，发现偏好建模也存在类似规律。研究收集多样化社区论坛数据，用1.5B-72B参数模型在1500万规模数据上进行训练，发现：(1)对抗性指标随数据和模型规模持续提升；(2)客观指标在大模型中呈现涌现特性；(3)主观指标无缩放趋势。实验表明WorldPM能显著提升不同规模偏好数据集(7K-800K样本)的泛化性能，关键子任务提升超5%。应用于强化学习人类反馈(RLHF)流程后，内部评估集提升4%-8%，验证了其作为偏好微调基础的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T20:01:35Z
- **目录日期**: 2025-05-17
