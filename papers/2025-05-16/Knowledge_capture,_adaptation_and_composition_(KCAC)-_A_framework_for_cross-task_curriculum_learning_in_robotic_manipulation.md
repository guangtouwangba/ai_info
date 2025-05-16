# Knowledge capture, adaptation and composition (KCAC): A framework for cross-task curriculum learning in robotic manipulation

**URL**: http://arxiv.org/abs/2505.10522v1

## 原始摘要

Reinforcement learning (RL) has demonstrated remarkable potential in robotic
manipulation but faces challenges in sample inefficiency and lack of
interpretability, limiting its applicability in real world scenarios. Enabling
the agent to gain a deeper understanding and adapt more efficiently to diverse
working scenarios is crucial, and strategic knowledge utilization is a key
factor in this process. This paper proposes a Knowledge Capture, Adaptation,
and Composition (KCAC) framework to systematically integrate knowledge transfer
into RL through cross-task curriculum learning. KCAC is evaluated using a two
block stacking task in the CausalWorld benchmark, a complex robotic
manipulation environment. To our knowledge, existing RL approaches fail to
solve this task effectively, reflecting deficiencies in knowledge capture. In
this work, we redesign the benchmark reward function by removing rigid
constraints and strict ordering, allowing the agent to maximize total rewards
concurrently and enabling flexible task completion. Furthermore, we define two
self-designed sub-tasks and implement a structured cross-task curriculum to
facilitate efficient learning. As a result, our KCAC approach achieves a 40
percent reduction in training time while improving task success rates by 10
percent compared to traditional RL methods. Through extensive evaluation, we
identify key curriculum design parameters subtask selection, transition timing,
and learning rate that optimize learning efficiency and provide conceptual
guidance for curriculum based RL frameworks. This work offers valuable insights
into curriculum design in RL and robotic learning.


## AI 摘要

这篇论文提出了一种名为KCAC（知识捕获、适应与组合）的强化学习框架，通过跨任务课程学习解决机器人操作中的样本低效和可解释性问题。研究在CausalWorld基准测试中使用双块堆叠任务验证了该框架的有效性：通过重新设计奖励函数（移除刚性约束和严格顺序）并设计两个子任务，KCAC比传统方法减少40%训练时间且成功率提升10%。研究还揭示了课程设计中的关键参数（子任务选择、过渡时机和学习率）对学习效率的影响，为基于课程的强化学习框架提供了理论指导。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T06:01:54Z
- **目录日期**: 2025-05-16
