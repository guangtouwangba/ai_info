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

该研究提出了一种新方法，通过模型融合而非微调来提升深度网络在小规模高价值数据（如医疗、安全领域）上的泛化能力。研究发现，只需对现有学习策略稍作调整，即可获得非平凡的泛化保证。关键突破在于：当采用数据驱动的下游任务融合学习时，认证泛化差距变得极小且与基础网络规模无关。实验首次在仅100个样本的情况下，为ViT-B和Mistral-7B等大型模型提供了有效的泛化证明。这一发现对现有系统的可信认证具有直接意义，并为实践与理论的交叉研究开辟了新方向。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T23:01:24Z
- **目录日期**: 2025-05-22
