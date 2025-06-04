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

随着AI技术进步，人类更倾向于将任务委托给AI代理。由于决策者通常无法完全了解影响委托决策的所有因素，他们会根据可观察特征对决策实例进行分类处理。研究提出了"类别化委托"问题，发现最优AI代理设计具有组合复杂性，且在一般情况下是计算困难的。不过，研究者在多个广泛场景中找到了高效算法，特别是当最优行动可分解为人类和算法各自观察特征的函数时。模拟实验表明，虽然迭代优化无法保证获得最优代理，但最终性能通常令人满意。这为设计人机协作算法提供了重要维度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-04T08:02:13Z
- **目录日期**: 2025-06-04
