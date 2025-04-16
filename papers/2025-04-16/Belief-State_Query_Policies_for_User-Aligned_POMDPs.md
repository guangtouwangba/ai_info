# Belief-State Query Policies for User-Aligned POMDPs

**URL**: http://arxiv.org/abs/2405.15907v2

## 原始摘要

Planning in real-world settings often entails addressing partial
observability while aligning with users' requirements. We present a novel
framework for expressing users' constraints and preferences about agent
behavior in a partially observable setting using parameterized belief-state
query (BSQ) policies in the setting of goal-oriented partially observable
Markov decision processes (gPOMDPs). We present the first formal analysis of
such constraints and prove that while the expected cost function of a
parameterized BSQ policy w.r.t its parameters is not convex, it is piecewise
constant and yields an implicit discrete parameter search space that is finite
for finite horizons. This theoretical result leads to novel algorithms that
optimize gPOMDP agent behavior with guaranteed user alignment. Analysis proves
that our algorithms converge to the optimal user-aligned behavior in the limit.
Empirical results show that parameterized BSQ policies provide a
computationally feasible approach for user-aligned planning in partially
observable settings.


## AI 摘要

本文提出了一种在部分可观测环境中（gPOMDPs框架下）通过参数化信念状态查询（BSQ）策略来表达用户约束与偏好的新方法。研究首次证明：虽然参数化BSQ策略的期望成本函数非凸，但具有分段常数特性，为有限时域问题生成隐式离散参数搜索空间。基于此理论突破，开发的新算法能保证优化智能体行为与用户需求对齐，并在极限情况下收敛至最优解。实验表明，参数化BSQ策略为部分可观测环境中的用户对齐规划提供了计算可行的解决方案。（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T17:03:20Z
- **目录日期**: 2025-04-16
