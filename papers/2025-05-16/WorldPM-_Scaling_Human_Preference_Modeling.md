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

该研究提出World Preference Modeling (WorldPM)，探索人类偏好的统一表示。通过收集公共论坛数据并使用1.5B到72B参数的模型在1500万规模数据上训练，发现：(1)对抗性指标随数据和模型规模提升；(2)客观指标在大模型中呈现涌现能力；(3)主观指标无规模效应。实验表明WorldPM能显著提升不同规模偏好数据集(7K-800K样本)的泛化性能，关键子任务提升超5%。集成到强化学习人类反馈(RLHF)流程后，内部和公开评估集分别获得4%-8%和显著提升，验证了其作为偏好微调基础的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T08:01:38Z
- **目录日期**: 2025-05-16
