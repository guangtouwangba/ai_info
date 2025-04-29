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

本文提出了一种可微概率层(PAL)，用于确保神经网络在连续环境中满足非凸代数约束，适用于自动驾驶等安全关键应用。PAL通过多项式参数化线性不等式组合，支持高效的符号积分归一化，可无缝集成到任何神经网络架构中，并通过最大似然训练。该方法实现了跨数据点的归一化计算分摊和GPU并行化，在代数约束集成基准测试和真实轨迹数据上验证了有效性。PAL特别擅长处理传统神经网络难以应对的复杂代数关系约束问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T12:02:29Z
- **目录日期**: 2025-04-29
