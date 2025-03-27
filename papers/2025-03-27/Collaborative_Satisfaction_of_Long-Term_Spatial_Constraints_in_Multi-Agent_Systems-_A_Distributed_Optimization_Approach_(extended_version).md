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

该研究提出了一种多智能体系统中协作满足长期空间约束的方法。每个智能体受到空间不等式约束的制约，这些约束可能取决于其他智能体的位置，且通信关系不确定。研究首先将问题建模为集中式无约束优化，通过最大化反映约束满足程度的目标函数获得最优配置，使智能体在满足自身约束的同时互相协作；当约束不可行时，系统收敛至最小违规解。随后提出了基于分布式共识的优化方案，近似集中式解，并开发了单积分器智能体的分布式控制器。仿真验证了该方法的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T01:28:53Z
- **目录日期**: 2025-03-27
