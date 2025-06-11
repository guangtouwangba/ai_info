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

本文研究发现，基于比较的偏好学习方法（如DPO、GPO和GBT）在调整AI模型与人类偏好时可能出现反直觉行为：即使偏好响应y胜过z，模型反而可能降低生成y的概率和奖励。作者通过理论分析证明，这类方法在温和假设下仍满足局部成对单调性，并提出多种单调性形式化定义及充分条件，为评估模型违反单调性的倾向提供了工具包。这些发现揭示了当前方法的局限性，为开发更可靠的偏好学习算法提供了指导。研究结果有助于理解现有方法的边界条件，推动更可信赖的AI对齐技术发展。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T07:02:20Z
- **目录日期**: 2025-06-11
