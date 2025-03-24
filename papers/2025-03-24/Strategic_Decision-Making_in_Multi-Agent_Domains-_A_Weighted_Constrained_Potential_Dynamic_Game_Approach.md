# Strategic Decision-Making in Multi-Agent Domains: A Weighted Constrained Potential Dynamic Game Approach

**URL**: http://arxiv.org/abs/2308.05876v3

## 原始摘要

In interactive multi-agent settings, decision-making and planning are
challenging mainly due to the agents' interconnected objectives. Dynamic game
theory offers a formal framework for analyzing such intricacies. Yet, solving
constrained dynamic games and determining the interaction outcome in the form
of generalized Nash Equilibria (GNE) pose computational challenges due to the
need for solving constrained coupled optimal control problems. In this paper,
we address this challenge by proposing to leverage the special structure of
many real-world multi-agent interactions. More specifically, our key idea is to
leverage constrained dynamic potential games, which are games for which GNE can
be found by solving a single constrained optimal control problem associated
with minimizing the potential function. We argue that constrained dynamic
potential games can effectively facilitate interactive decision-making in many
multi-agent interactions. We will identify structures in realistic multi-agent
interactive scenarios that can be transformed into weighted constrained
potential dynamic games (WCPDGs). We will show that the GNE of the resulting
WCPDG can be obtained by solving a single constrained optimal control problem.
We will demonstrate the effectiveness of the proposed method through various
simulation studies and show that we achieve significant improvements in solve
time compared to state-of-the-art game solvers. We further provide experimental
validation of our proposed method in a navigation setup involving two
quadrotors carrying a rigid object while avoiding collisions with two humans.


## AI 摘要

本文提出利用约束动态势博弈（WCPDG）来解决多智能体交互决策中的计算挑战。通过分析现实场景中的特殊结构，将问题转化为加权约束动态势博弈，其广义纳什均衡（GNE）可通过求解单个约束最优控制问题获得。相比现有博弈求解器，该方法显著缩短了求解时间。仿真实验验证了其有效性，并在双四旋翼无人机协同搬运物体并避让行人的导航任务中进行了实际验证。该方法通过势函数最小化简化了复杂耦合优化问题，为多目标交互决策提供了高效解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-24T21:02:03Z
- **目录日期**: 2025-03-24
