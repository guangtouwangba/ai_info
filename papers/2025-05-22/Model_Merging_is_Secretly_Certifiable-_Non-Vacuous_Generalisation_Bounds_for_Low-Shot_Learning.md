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

该研究首次证明了深度网络在小样本高风险任务中的非平凡泛化能力。通过将模型融合方法与泛化认证理论相结合，研究发现只需对现有学习策略稍作调整，即可获得与基础网络规模无关的微小泛化误差保证。实验表明，仅需100个样本，就能为VIT-B视觉模型和mistral-7B语言模型提供可验证的泛化保证。这一发现对医疗、安全等关键领域的AI系统认证具有重要意义，为实践与理论的交叉研究开辟了新方向，使现有系统更容易获得可信认证。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T06:01:30Z
- **目录日期**: 2025-05-22
