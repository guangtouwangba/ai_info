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

这篇论文研究了基于比较的偏好学习方法（如DPO、GPO和GBT）中出现的非单调性问题，即模型有时会降低被偏好响应的概率。作者通过实证观察发现这一现象后，在温和假设下证明了这类方法仍满足局部成对单调性，并提出了多种单调性形式化定义及保证条件。研究结果为理解现有方法的局限性提供了理论工具，并为开发更可靠的偏好学习算法提供了指导，有助于改善AI模型与人类偏好的对齐效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-12T00:02:11Z
- **目录日期**: 2025-06-12
