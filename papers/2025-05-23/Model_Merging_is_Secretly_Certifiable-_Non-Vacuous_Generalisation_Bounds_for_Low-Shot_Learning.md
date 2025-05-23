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

这篇论文探讨了在医学、安全等高风险领域，如何为深度网络提供非空泛的IID泛化能力证明。研究发现，通过将模型融合方法与泛化证明相结合，并对现有学习策略稍作调整，即可获得非平凡的泛化保证。关键在于采用数据驱动的下游任务融合学习而非微调，这使得泛化差距极小且与基础网络规模无关。实验首次展示了在仅用100个样本的情况下，VIT-B视觉模型和mistral-7B语言模型仍能保持非平凡泛化性能。这一发现对现有系统的可信认证具有直接意义，并为实践与理论的交叉研究开辟了新方向。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T00:01:36Z
- **目录日期**: 2025-05-23
