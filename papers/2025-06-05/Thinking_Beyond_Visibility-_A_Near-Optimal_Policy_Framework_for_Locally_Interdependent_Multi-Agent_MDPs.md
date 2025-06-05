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

分散式部分可观察马尔可夫决策过程(Dec-POMDPs)通常难以求解。DeWeese和Qu(2024)针对局部可见性和依赖性问题提出了局部相互依赖多智能体MDP模型，并给出了三种可计算且接近最优的策略。但这些策略在小可见度下会出现"惩罚抖动"现象。本文提出了扩展截断策略类，这是首个针对任意局部相互依赖多智能体MDP的非平凡近优闭式策略类，能记忆超出可见范围的智能体，显著改善小可见度性能，解决惩罚抖动问题，并在特定条件下保证完全可观察的最优行为。同时扩展了原模型，使其支持转移依赖和扩展奖励依赖。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T17:01:16Z
- **目录日期**: 2025-06-05
