# MultiNash-PF: A Particle Filtering Approach for Computing Multiple Local Generalized Nash Equilibria in Trajectory Games

**URL**: http://arxiv.org/abs/2410.05554v2

## 原始摘要

Modern robotic systems frequently engage in complex multi-agent interactions,
many of which are inherently multi-modal, meaning they can lead to multiple
distinct outcomes. To interact effectively, robots must recognize the possible
interaction modes and adapt to the one preferred by other agents. In this work,
we propose an efficient algorithm for capturing the multimodality in
multi-agent interactions. We leverage a game-theoretic planner to model
interaction outcomes as equilibria where \emph{each equilibrium} corresponds to
a distinct interaction \emph{mode}. We then develop an efficient algorithm to
identify all the equilibria, allowing robots to reason about multiple
interaction modes. More specifically, we formulate interactive planning as
Constrained Potential Trajectory Games (CPTGs) and model interaction outcomes
by local Generalized Nash equilibria (GNEs) of the game. CPTGs are a class of
games for which a local GNE can be found by solving a single constrained
optimal control problem where a potential function is minimized. We propose to
integrate the potential game approach with implicit particle filtering, a
sample-efficient method for non-convex trajectory optimization. We utilize
implicit particle filtering to identify the coarse estimates of multiple local
minimizers of the game's potential function. MultiNash-PF then refines these
estimates with optimization solvers, obtaining different local GNEs. We show
through numerical simulations that MultiNash-PF reduces computation time by up
to 50\% compared to a baseline. We further demonstrate the effectiveness of our
algorithm in real-world human-robot interaction scenarios, where it
successfully accounts for the multi-modal nature of interactions and resolves
potential conflicts in real-time.


## AI 摘要

本文提出了一种高效算法MultiNash-PF，用于捕捉多智能体交互中的多模态特性。通过将交互规划建模为约束势能轨迹游戏（CPTGs），并利用局部广义纳什均衡（GNEs）表示交互结果，该算法结合了势能游戏方法和隐式粒子滤波技术，有效识别多个局部极小值点。MultiNash-PF通过优化求解器细化这些估计，获得不同的局部GNEs。数值模拟显示，该算法比基线方法减少高达50%的计算时间，并在真实人机交互场景中成功应对多模态交互并实时解决潜在冲突。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T05:02:30Z
- **目录日期**: 2025-03-24
