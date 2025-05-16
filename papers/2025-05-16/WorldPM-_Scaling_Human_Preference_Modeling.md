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

研究发现偏好建模（WorldPM）与语言模型类似，存在训练数据量和模型规模扩展的幂律关系。基于15M数据对1.5B-72B参数模型的实验显示：对抗性指标随规模持续提升；客观指标在大模型中呈现涌现能力；主观指标无扩展规律。WorldPM作为偏好微调基础，在7个基准测试（20个子任务）中显著提升泛化性能，关键任务改进超5%。应用于RLHF流程后，内部和公共评估集分别获得4%-8%和显著提升。该研究验证了统一人类偏好表征的扩展潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T18:01:27Z
- **目录日期**: 2025-05-16
