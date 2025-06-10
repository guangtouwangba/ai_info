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

当前AI测试时扩展方法主要依赖生成长推理轨迹（"多思考"）再做出响应，但在需要交互的任务中无法动态获取环境信息。本研究提出测试时交互扩展（TTI），通过延长交互时间实现探索、回溯和动态重规划等行为。针对网页代理任务，研究发现仅通过提示式交互扩展就能显著提升性能。进一步提出基于课程学习的在线强化学习TTI方法，自适应调整交互长度。使用Gemma 3 12B模型的TTI在WebVoyager和WebArena基准上创造了开源最佳成绩，能自适应平衡探索与利用，证明交互扩展是计算扩展的有力补充维度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T10:01:58Z
- **目录日期**: 2025-06-10
