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

该研究提出**World Preference Modeling (WorldPM)**，探索偏好建模中的规模定律，发现其与语言模型类似，测试损失随模型和数据规模呈幂律关系。基于15M规模的多样化论坛数据，实验覆盖1.5B至72B参数模型，结果显示：(1) 对抗性指标（识别欺骗特征）随规模稳定提升；(2) 客观指标（明确答案任务）在大模型中涌现能力；(3) 主观指标（有限人类或AI偏好）无规模效应。WorldPM作为偏好微调基础，在7个基准测试的20项子任务中普遍提升泛化性能（关键任务增益超5%），集成至RLHF流程后，内部评估提升4%-8%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T15:01:45Z
- **目录日期**: 2025-05-18
