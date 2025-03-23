# Computing Lindahl Equilibrium for Public Goods with and without Funding Caps

**URL**: http://arxiv.org/abs/2503.16414v1

## 原始摘要

Lindahl equilibrium is a solution concept for allocating a fixed budget
across several divisible public goods. It always lies in the core, meaning that
the equilibrium allocation satisfies desirable stability and proportional
fairness properties. We consider a model where agents have separable linear
utility functions over the public goods, and the output assigns to each good an
amount of spending, summing to at most the available budget.
  In the uncapped setting, each of the public goods can absorb any amount of
funding. In this case, it is known that Lindahl equilibrium is equivalent to
maximizing Nash social welfare, and this allocation can be computed by a
public-goods variant of the proportional response dynamics. We introduce a new
convex programming formulation for computing this solution and show that it is
related to Nash welfare maximization through duality and reformulation. We then
show that the proportional response dynamics is equivalent to running mirror
descent on our new formulation, thereby providing a new and immediate proof of
the convergence guarantee for the dynamics. Our new formulation has
similarities to Shmyrev's convex program for Fisher market equilibrium.
  In the capped setting, each public good has an upper bound on the amount of
funding it can receive. In this setting, existence of Lindahl equilibrium was
only known via fixed-point arguments. The existence of an efficient algorithm
computing one has been a long-standing open question. We prove that our new
convex program continues to work when the cap constraints are added, and its
optimal solutions are Lindahl equilibria. Thus, we establish that Lindahl
equilibrium can be efficiently computed in the capped setting. Our result also
implies that approximately core-stable allocations can be efficiently computed
for the class of separable piecewise-linear concave (SPLC) utilities.


## AI 摘要

Lindahl均衡是一种用于在多个可分公共物品之间分配固定预算的解决方案，具有稳定性和比例公平性。本文研究了在无上限和有上限预算约束下，Lindahl均衡的计算问题。在无上限情况下，Lindahl均衡等价于最大化Nash社会福利，并提出了一种新的凸规划公式，通过对偶性和重构与Nash福利最大化相关联。在有上限情况下，证明了该凸规划公式仍适用，且其最优解为Lindahl均衡，从而解决了长期存在的计算问题。此外，结果表明，对于可分分段线性凹效用函数，近似核心稳定分配可以高效计算。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-23T15:01:23Z
- **目录日期**: 2025-03-23
