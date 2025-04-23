# RILe: Reinforced Imitation Learning

**URL**: http://arxiv.org/abs/2406.08472v4

## 原始摘要

Acquiring complex behaviors is essential for artificially intelligent agents,
yet learning these behaviors in high-dimensional settings poses a significant
challenge due to the vast search space. Traditional reinforcement learning (RL)
requires extensive manual effort for reward function engineering. Inverse
reinforcement learning (IRL) uncovers reward functions from expert
demonstrations but relies on an iterative process that is often computationally
expensive. Imitation learning (IL) provides a more efficient alternative by
directly comparing an agent's actions to expert demonstrations; however, in
high-dimensional environments, such direct comparisons often offer insufficient
feedback for effective learning. We introduce RILe (Reinforced Imitation
Learning), a framework that combines the strengths of imitation learning and
inverse reinforcement learning to learn a dense reward function efficiently and
achieve strong performance in high-dimensional tasks. RILe employs a novel
trainer-student framework: the trainer learns an adaptive reward function, and
the student uses this reward signal to imitate expert behaviors. By dynamically
adjusting its guidance as the student evolves, the trainer provides nuanced
feedback across different phases of learning. Our framework produces
high-performing policies in high-dimensional tasks where direct imitation fails
to replicate complex behaviors. We validate RILe in challenging robotic
locomotion tasks, demonstrating that it significantly outperforms existing
methods and achieves near-expert performance across multiple settings.


## AI 摘要

本文提出RILe（强化模仿学习）框架，通过结合模仿学习(IL)和逆向强化学习(IRL)的优势，解决高维环境中复杂行为学习的挑战。RILe采用"训练者-学习者"架构：训练者学习自适应奖励函数，学习者利用该奖励信号模仿专家行为。该方法能动态调整指导策略，在不同学习阶段提供精细反馈。实验证明，RILe在机器人运动等高维任务中显著优于现有方法，几乎达到专家水平，而直接模仿学习在这些复杂任务中往往失效。该框架避免了传统RL需要人工设计奖励函数的问题，同时比IRL更高效。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-23T02:30:18Z
- **目录日期**: 2025-04-23
