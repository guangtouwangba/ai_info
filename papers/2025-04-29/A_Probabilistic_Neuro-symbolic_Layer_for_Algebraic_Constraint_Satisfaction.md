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

这篇论文提出了一种可微的概率代数层(PAL)，用于确保神经网络在连续环境中满足非凸代数约束。PAL可以无缝集成到任何神经架构中，通过最大似然进行训练，无需近似处理。它定义了基于多项式参数化的线性不等式组合的概率分布，支持高效的符号积分精确归一化，并可在GPU上并行计算。该方法适用于自动驾驶等安全关键场景，如避免碰撞或偏离道路。实验验证了PAL在代数约束集成和真实轨迹数据上的有效性。关键创新在于提供了可微、可训练且严格满足复杂约束的神经网络层。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T14:02:17Z
- **目录日期**: 2025-04-29
