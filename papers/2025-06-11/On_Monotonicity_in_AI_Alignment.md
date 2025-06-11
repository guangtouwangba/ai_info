# On Monotonicity in AI Alignment

**URL**: http://arxiv.org/abs/2506.08998v1

## 原始摘要

Comparison-based preference learning has become central to the alignment of
AI models with human preferences. However, these methods may behave
counterintuitively. After empirically observing that, when accounting for a
preference for response $y$ over $z$, the model may actually decrease the
probability (and reward) of generating $y$ (an observation also made by
others), this paper investigates the root causes of (non) monotonicity, for a
general comparison-based preference learning framework that subsumes Direct
Preference Optimization (DPO), Generalized Preference Optimization (GPO) and
Generalized Bradley-Terry (GBT). Under mild assumptions, we prove that such
methods still satisfy what we call local pairwise monotonicity. We also provide
a bouquet of formalizations of monotonicity, and identify sufficient conditions
for their guarantee, thereby providing a toolbox to evaluate how prone learning
models are to monotonicity violations. These results clarify the limitations of
current methods and provide guidance for developing more trustworthy preference
learning algorithms.


## AI 摘要

这篇论文研究了基于比较的偏好学习方法（如DPO、GPO和GBT）中出现的反直觉现象：即使模型学习到偏好响应y胜过z，却可能降低生成y的概率。作者通过理论分析证明，这些方法仍满足局部成对单调性，并提出了一套形式化框架来评估不同单调性条件的满足情况。研究揭示了当前方法的局限性，为开发更可靠的偏好学习算法提供了理论指导。这些发现有助于理解AI模型与人类偏好对齐过程中的潜在问题，并为改进算法设计提供了工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T09:02:21Z
- **目录日期**: 2025-06-11
