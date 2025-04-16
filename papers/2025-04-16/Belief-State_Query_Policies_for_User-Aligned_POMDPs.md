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

本文提出了一种在部分可观测环境中实现用户对齐规划的新框架，采用参数化信念状态查询(BSQ)策略在目标导向POMDP(gPOMDP)中表达用户约束和偏好。理论分析表明，参数化BSQ策略的期望成本函数虽非凸，但分段恒定，隐含有限离散参数空间。基于此，研究者开发了能保证用户对齐的新型算法，理论证明其可渐近收敛至最优用户对齐行为。实验验证了参数化BSQ策略在部分可观测环境下进行用户对齐规划的可行性。该研究为复杂环境中的智能体行为优化提供了理论保证和实用方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-16T21:03:06Z
- **目录日期**: 2025-04-16
