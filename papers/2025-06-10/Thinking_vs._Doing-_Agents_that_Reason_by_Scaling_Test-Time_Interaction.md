# Thinking vs. Doing: Agents that Reason by Scaling Test-Time Interaction

**URL**: http://arxiv.org/abs/2506.07976v1

## 原始摘要

The current paradigm of test-time scaling relies on generating long reasoning
traces ("thinking" more) before producing a response. In agent problems that
require interaction, this can be done by generating thinking traces before
acting in the world. However, this process does not allow agents to acquire new
information from the environment or adapt their behavior over time. In this
work, we propose to scale test-time interaction, an untapped dimension of
test-time scaling that increases the agent's interaction horizon to enable
running rich behaviors such as exploration, backtracking, and dynamic
re-planning within a single rollout. To demonstrate the promise of this scaling
dimension, we study the domain of web agents. We first show that even
prompting-based interaction scaling without any training can improve task
success on web benchmarks non-trivially. Building on this, we introduce TTI
(Test-Time Interaction), a curriculum-based online reinforcement learning (RL)
approach that trains agents by adaptively adjusting their rollout lengths.
Using a Gemma 3 12B model, TTI produces state-of-the-art open-source, open-data
web agents on WebVoyager and WebArena benchmarks. We further show that TTI
enables agents to balance exploration and exploitation adaptively. Our results
establish interaction scaling as a powerful, complementary axis to scaling
per-step compute, offering new avenues for training adaptive agents.


## AI 摘要

本文提出了一种新的测试时扩展维度——交互扩展（TTI），通过延长智能体与环境的交互时间来实现更复杂的行为（如探索、回溯和动态重规划）。在网页智能体任务中，实验表明仅通过提示驱动的交互扩展就能显著提升任务成功率。进一步提出的TTI方法采用课程式在线强化学习，动态调整智能体行动轨迹长度。基于Gemma 3 12B模型的TTI在WebVoyager和WebArena基准测试中创造了开源网页智能体的新记录，并能自适应平衡探索与利用。该研究证明交互扩展是独立于单步计算扩展的有效扩展维度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T14:01:56Z
- **目录日期**: 2025-06-10
