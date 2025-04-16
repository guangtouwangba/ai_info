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

本文提出了一种在部分可观测环境中（gPOMDPs框架下）通过参数化信念状态查询（BSQ）策略来表达用户约束和偏好的新方法。理论分析表明，虽然参数化BSQ策略的期望成本函数非凸，但具有分段常数特性，并隐含有限的离散参数搜索空间。基于此，研究开发了能保证用户对齐的新算法，理论上可收敛至最优用户对齐行为。实验验证了该方法在部分可观测环境下实现用户对齐规划的可行性。该框架首次为部分可观测环境中的用户需求对齐提供了形式化分析和解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T22:03:03Z
- **目录日期**: 2025-04-16
