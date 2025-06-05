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

分散式部分可观察马尔可夫决策过程（Dec-POMDPs）通常难以求解，但局部可见性和依赖性的假设可简化问题。DeWeese和Qu（2024）提出局部依赖多智能体MDP框架，并给出三种易计算且接近最优的闭式策略，但在有限可见度下可能因"惩罚抖动"现象表现不佳。本研究提出扩展截断策略类，首次实现任意局部依赖多智能体MDP中接近最优的闭式部分可观察策略，通过记忆超出可见范围的智能体显著提升性能，解决抖动问题，并在特定条件下实现完全可观察的最优行为。同时扩展了模型以支持转移依赖和广义奖励依赖。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-05T20:01:08Z
- **目录日期**: 2025-06-05
