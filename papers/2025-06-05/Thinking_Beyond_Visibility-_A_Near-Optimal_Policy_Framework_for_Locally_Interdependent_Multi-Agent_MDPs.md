# Thinking Beyond Visibility: A Near-Optimal Policy Framework for Locally Interdependent Multi-Agent MDPs

**URL**: http://arxiv.org/abs/2506.04215v1

## 原始摘要

Decentralized Partially Observable Markov Decision Processes (Dec-POMDPs) are
known to be NEXP-Complete and intractable to solve. However, for problems such
as cooperative navigation, obstacle avoidance, and formation control, basic
assumptions can be made about local visibility and local dependencies. The work
DeWeese and Qu 2024 formalized these assumptions in the construction of the
Locally Interdependent Multi-Agent MDP. In this setting, it establishes three
closed-form policies that are tractable to compute in various situations and
are exponentially close to optimal with respect to visibility. However, it is
also shown that these solutions can have poor performance when the visibility
is small and fixed, often getting stuck during simulations due to the so called
"Penalty Jittering" phenomenon. In this work, we establish the Extended Cutoff
Policy Class which is, to the best of our knowledge, the first non-trivial
class of near optimal closed-form partially observable policies that are
exponentially close to optimal with respect to the visibility for any Locally
Interdependent Multi-Agent MDP. These policies are able to remember agents
beyond their visibilities which allows them to perform significantly better in
many small and fixed visibility settings, resolve Penalty Jittering
occurrences, and under certain circumstances guarantee fully observable joint
optimal behavior despite the partial observability. We also propose a
generalized form of the Locally Interdependent Multi-Agent MDP that allows for
transition dependence and extended reward dependence, then replicate our
theoretical results in this setting.


## AI 摘要

本文介绍了针对局部可观测多智能体马尔可夫决策过程(Locally Interdependent Multi-Agent MDP)提出的扩展截断策略类(Extended Cutoff Policy Class)。该策略类能在局部可观测条件下，通过记忆超出可视范围的智能体信息，显著提升小固定可视范围场景的性能，解决"惩罚抖动"问题，并在特定条件下实现完全可观测的最优行为。研究还扩展了原有模型框架，允许转移依赖和扩展奖励依赖，并在新框架下复现了理论结果。这些策略计算可行且随可视范围呈指数级接近最优，是首个非平凡的最优闭式部分可观测策略类。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T10:00:54Z
- **目录日期**: 2025-06-05
