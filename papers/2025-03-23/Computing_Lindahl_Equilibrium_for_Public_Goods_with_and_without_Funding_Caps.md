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

林达尔均衡是一种用于在多个可分公共物品之间分配固定预算的解决方案概念，始终位于核心，满足稳定性和比例公平性。本文研究了代理人对公共物品具有可分离线性效用函数的模型，并引入了一种新的凸规划公式来计算该解，证明了其与纳什福利最大化相关。通过镜像下降法，比例响应动态的收敛性得到了新的证明。在有限制条件下，本文证明了新的凸规划公式仍然有效，其最优解为林达尔均衡，从而解决了长期存在的算法效率问题，并表明可有效计算近似核心稳定分配。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-23T13:04:55Z
- **目录日期**: 2025-03-23
