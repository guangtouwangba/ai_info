# Thinking vs. Doing: Agents that Reason by Scaling Test-Time Interaction

**URL**: http://arxiv.org/abs/2506.07976v2

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

本文提出了一种新的测试时交互扩展方法（TTI），通过延长智能体的交互时间来实现更复杂的行为（如探索、回溯和动态重规划）。不同于传统仅增加推理步骤的测试时扩展，TTI允许智能体在单次运行中从环境获取新信息并动态调整策略。研究以网页智能体为案例，发现仅通过提示驱动的交互扩展就能显著提升任务成功率。进一步提出的TTI方法采用课程式在线强化学习，自适应调整交互时长，基于Gemma 3 12B模型在WebVoyager和WebArena基准上达到开源最佳性能，证明了交互扩展是计算扩展的有效补充维度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T02:31:58Z
- **目录日期**: 2025-06-11
