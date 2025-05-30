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

该研究提出了一种高效算法MultiNash-PF，用于处理多智能体交互中的多模态问题。通过将交互建模为约束势轨迹博弈（CPTG），将不同交互模式对应为博弈的广义纳什均衡（GNE）。算法结合势博弈与隐式粒子滤波，先粗定位势函数的多个局部极小值，再通过优化求解器精修得到不同GNE。实验显示，相比基线方法计算时间减少50%，并在人机交互场景中有效识别多模态交互、实时解决潜在冲突。该方法为机器人适应复杂多智能体交互提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T23:01:52Z
- **目录日期**: 2025-03-24
