# Designing Algorithmic Delegates: The Role of Indistinguishability in Human-AI Handoff

**URL**: http://arxiv.org/abs/2506.03102v1

## 原始摘要

As AI technologies improve, people are increasingly willing to delegate tasks
to AI agents. In many cases, the human decision-maker chooses whether to
delegate to an AI agent based on properties of the specific instance of the
decision-making problem they are facing. Since humans typically lack full
awareness of all the factors relevant to this choice for a given
decision-making instance, they perform a kind of categorization by treating
indistinguishable instances -- those that have the same observable features --
as the same. In this paper, we define the problem of designing the optimal
algorithmic delegate in the presence of categories. This is an important
dimension in the design of algorithms to work with humans, since we show that
the optimal delegate can be an arbitrarily better teammate than the optimal
standalone algorithmic agent. The solution to this optimal delegation problem
is not obvious: we discover that this problem is fundamentally combinatorial,
and illustrate the complex relationship between the optimal design and the
properties of the decision-making task even in simple settings. Indeed, we show
that finding the optimal delegate is computationally hard in general. However,
we are able to find efficient algorithms for producing the optimal delegate in
several broad cases of the problem, including when the optimal action may be
decomposed into functions of features observed by the human and the algorithm.
Finally, we run computational experiments to simulate a designer updating an
algorithmic delegate over time to be optimized for when it is actually adopted
by users, and show that while this process does not recover the optimal
delegate in general, the resulting delegate often performs quite well.


## AI 摘要

随着AI技术发展，人类更倾向于将任务委托给AI代理。由于人类通常无法完全掌握决策所需的所有因素，他们会基于可观察特征对相似决策实例进行归类处理。研究发现，在分类场景下设计最优算法代理具有重要价值——这种委托代理可能显著优于独立算法代理。该问题本质上是组合优化问题，最优设计与决策任务特性存在复杂关联，且通常属于计算难题。不过在某些场景（如最优行动可分解为人类与算法各自观察特征的函数时）存在高效解法。实验表明，通过迭代更新设计的代理虽非全局最优，但通常表现良好。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T18:02:01Z
- **目录日期**: 2025-06-04
