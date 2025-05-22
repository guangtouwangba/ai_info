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

这篇论文提出了一种新方法，通过模型融合而非微调来认证深度网络在独立同分布(IID)数据上的泛化能力。研究发现，只需对现有学习策略稍作调整，就能为小规模数据(低至100个样本)提供非平凡的泛化保证，且保证与基础网络规模无关。该方法在视觉模型(VIT-B)和语言模型(mistral-7B)上都验证了有效性。这一发现具有重要意义，既为现有系统提供了可信认证的可行路径，也为理论与实践的结合开辟了新研究方向，有助于推动AI在医疗、安全等高风险领域的可信应用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T11:01:34Z
- **目录日期**: 2025-05-22
