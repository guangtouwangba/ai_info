# A Probabilistic Neuro-symbolic Layer for Algebraic Constraint Satisfaction

**URL**: http://arxiv.org/abs/2503.19466v2

## 原始摘要

In safety-critical applications, guaranteeing the satisfaction of constraints
over continuous environments is crucial, e.g., an autonomous agent should never
crash into obstacles or go off-road. Neural models struggle in the presence of
these constraints, especially when they involve intricate algebraic
relationships. To address this, we introduce a differentiable probabilistic
layer that guarantees the satisfaction of non-convex algebraic constraints over
continuous variables. This probabilistic algebraic layer (PAL) can be
seamlessly plugged into any neural architecture and trained via maximum
likelihood without requiring approximations. PAL defines a distribution over
conjunctions and disjunctions of linear inequalities, parameterized by
polynomials. This formulation enables efficient and exact renormalization via
symbolic integration, which can be amortized across different data points and
easily parallelized on a GPU. We showcase PAL and our integration scheme on a
number of benchmarks for algebraic constraint integration and on real-world
trajectory data.


## AI 摘要

在安全关键应用中，确保连续环境下的约束条件至关重要（如自动驾驶需规避障碍物）。针对神经网络难以处理复杂代数约束的问题，本研究提出可微分概率代数层（PAL），通过多项式参数化线性不等式组合，实现非凸代数约束的严格满足。PAL可无缝嵌入任何神经架构，支持精确符号积分归一化，并支持GPU并行加速。实验验证了其在代数约束集成和真实轨迹数据上的有效性，为安全关键系统提供了可训练的约束保障方案。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T19:02:03Z
- **目录日期**: 2025-04-29
