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

研究发现偏好建模（WorldPM）也存在类似语言模型的规模法则，即性能随模型和数据规模呈幂律增长。通过收集论坛多样用户偏好数据，在1.5B到72B参数模型上训练1500万规模数据发现：(1)对抗性指标（识别欺骗特征）随规模稳定提升；(2)客观指标（明确答案）在大模型中呈现涌现能力；(3)主观指标无规模效应。实验表明WorldPM作为偏好微调基础能广泛提升不同规模数据集（7K-800K样本）的泛化性能，关键子任务提升超5%。应用于强化学习人类反馈（RLHF）流程后，内部评估集提升4-8%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T07:01:23Z
- **目录日期**: 2025-05-17
