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

本文提出了一种可微分概率层(PAL)，用于确保神经网络在连续环境中满足非凸代数约束。PAL可无缝集成到任何神经架构中，通过最大似然进行训练，无需近似处理。它定义了基于多项式参数化的线性不等式合取和析取分布，支持通过符号积分进行高效精确的归一化，并可在GPU上并行处理。该方法在代数约束集成基准测试和真实轨迹数据上进行了验证，适用于自动驾驶等安全关键场景，确保智能体不会碰撞或偏离道路。PAL解决了神经网络在处理复杂代数关系时的约束满足难题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T23:01:59Z
- **目录日期**: 2025-04-29
