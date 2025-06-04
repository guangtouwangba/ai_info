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

随着AI技术进步，人们更倾向于将任务委托给AI代理。人类决策者通常基于可观察特征对问题实例进行分类，并决定是否委托AI处理。研究发现，在这种分类情境下设计最优算法代理至关重要，其表现可能远超独立算法代理。但最优委托问题具有组合复杂性，计算难度高。研究提出了在某些场景（如人类与算法特征可分解时）的高效求解方案，并通过模拟实验证明：虽然迭代更新不一定获得理论最优解，但实际性能往往表现良好。这为AI与人类协作系统的设计提供了重要洞见。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T03:22:35Z
- **目录日期**: 2025-06-04
