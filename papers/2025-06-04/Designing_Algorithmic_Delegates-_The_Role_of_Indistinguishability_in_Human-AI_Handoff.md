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

随着AI技术的进步，人们更倾向于将任务委托给AI代理。人类决策者通常基于可观察特征对决策问题进行分类，选择是否委托AI处理。研究发现，在分类情境下设计最优AI代理具有重要价值，其表现可能远超独立算法代理。该问题具有组合复杂性，最优方案与任务特性密切相关，且通常属于计算难题。不过，在某些场景（如最优行动可分解为人类与算法观察特征的函数时）存在高效解法。模拟实验表明，通过迭代优化的代理虽非全局最优，但能实现良好性能。这一发现为算法设计提供了新视角。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T05:02:09Z
- **目录日期**: 2025-06-04
