# Model Merging is Secretly Certifiable: Non-Vacuous Generalisation Bounds for Low-Shot Learning

**URL**: http://arxiv.org/abs/2505.15798v1

## 原始摘要

Certifying the IID generalisation ability of deep networks is the first of
many requirements for trusting AI in high-stakes applications from medicine to
security. However, when instantiating generalisation bounds for deep networks
it remains challenging to obtain non-vacuous guarantees, especially when
applying contemporary large models on the small scale data prevalent in such
high-stakes fields. In this paper, we draw a novel connection between a family
of learning methods based on model fusion and generalisation certificates, and
surprisingly show that with minor adjustment several existing learning
strategies already provide non-trivial generalisation guarantees. Essentially,
by focusing on data-driven learning of downstream tasks by fusion rather than
fine-tuning, the certified generalisation gap becomes tiny and independent of
the base network size, facilitating its certification. Our results show for the
first time non-trivial generalisation guarantees for learning with as low as
100 examples, while using vision models such as VIT-B and language models such
as mistral-7B. This observation is significant as it has immediate implications
for facilitating the certification of existing systems as trustworthy, and
opens up new directions for research at the intersection of practice and
theory.


## AI 摘要

这篇论文提出了一种新方法，通过模型融合而非微调来认证深度网络在独立同分布(IID)数据上的泛化能力。研究发现，只需对现有学习策略稍作调整，就能为小规模数据（低至100样本）提供非平凡的泛化保证，且该保证与基础网络规模无关。实验证明，该方法适用于视觉模型（如VIT-B）和语言模型（如mistral-7B），首次实现了在小样本场景下的可认证泛化。这一发现对医疗、安全等高风险领域AI系统的可信认证具有直接意义，并为理论与实践的结合开辟了新方向。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T15:01:35Z
- **目录日期**: 2025-05-22
