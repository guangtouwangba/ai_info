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

Dec-POMDPs通常难以求解，但针对局部可见性和依赖性问题（如协作导航），DeWeese和Qu（2024）提出了局部依赖多智能体MDP框架，给出了三种可计算且接近最优的闭式策略。然而，这些策略在小可见度下可能因"惩罚抖动"现象表现不佳。本文提出扩展截断策略类，首次实现了对任意局部依赖多智能体MDP的接近最优闭式部分可观测策略，能记忆超出可见范围的智能体，显著提升小可见度下的性能，消除抖动现象，并在特定条件下实现完全可观测的最优行为。研究还扩展了该框架以支持转移依赖和广义奖励依赖。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T07:01:25Z
- **目录日期**: 2025-06-05
