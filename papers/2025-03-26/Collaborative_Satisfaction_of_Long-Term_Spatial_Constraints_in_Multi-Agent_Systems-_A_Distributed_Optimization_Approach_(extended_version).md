# Collaborative Satisfaction of Long-Term Spatial Constraints in Multi-Agent Systems: A Distributed Optimization Approach (extended version)

**URL**: http://arxiv.org/abs/2503.19879v1

## 原始摘要

This paper addresses the problem of collaboratively satisfying long-term
spatial constraints in multi-agent systems. Each agent is subject to spatial
constraints, expressed as inequalities, which may depend on the positions of
other agents with whom they may or may not have direct communication. These
constraints need to be satisfied asymptotically or after an unknown finite
time. The agents' objective is to collectively achieve a formation that
fulfills all constraints. The problem is initially framed as a centralized
unconstrained optimization, where the solution yields the optimal configuration
by maximizing an objective function that reflects the degree of constraint
satisfaction. This function encourages collaboration, ensuring agents help each
other meet their constraints while fulfilling their own. When the constraints
are infeasible, agents converge to a least-violating solution. A distributed
consensus-based optimization scheme is then introduced, which approximates the
centralized solution, leading to the development of distributed controllers for
single-integrator agents. Finally, simulations validate the effectiveness of
the proposed approach.


## AI 摘要

该研究提出了一种多智能体系统中协作满足长期空间约束的方法。每个智能体都受到空间不等式约束，这些约束可能依赖于其他智能体的位置，无论是否存在直接通信。系统通过集中式无约束优化框架，最大化反映约束满足程度的协作目标函数，使智能体在满足自身约束的同时协助其他智能体。当约束不可行时，系统会收敛到最小违规解。研究还提出了基于分布式共识的优化方案，近似集中式解，并开发了单积分器智能体的分布式控制器。仿真验证了该方法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T16:01:55Z
- **目录日期**: 2025-03-26
