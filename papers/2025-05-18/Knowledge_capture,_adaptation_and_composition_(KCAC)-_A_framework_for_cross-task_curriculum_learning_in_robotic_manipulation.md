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

本文提出了一种知识捕获、适应与组合（KCAC）框架，通过跨任务课程学习将知识迁移系统性地融入强化学习（RL），以解决机器人操作中样本效率低和可解释性差的问题。在CausalWorld基准测试中，KCAC重构了奖励函数，设计子任务并采用结构化课程学习，使训练时间减少40%，任务成功率提升10%。研究还确定了优化学习效率的关键课程设计参数（子任务选择、过渡时机和学习率），为基于课程的RL框架提供了理论指导。该成果为RL和机器人学习的课程设计提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-18T23:01:45Z
- **目录日期**: 2025-05-18
