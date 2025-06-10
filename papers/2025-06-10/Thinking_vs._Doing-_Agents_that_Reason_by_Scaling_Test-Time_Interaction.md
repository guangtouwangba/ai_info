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

当前测试时扩展范式主要依赖生成长推理轨迹（"多思考"）再行动，但在交互式任务中无法动态获取环境信息或调整行为。本研究提出测试时交互扩展这一新维度，通过延长交互时间实现探索、回溯和动态重规划等复杂行为。以网页代理为例，实验表明仅通过提示式交互扩展即可显著提升任务成功率。进一步提出的TTI方法（基于课程学习的在线强化学习）能自适应调整交互长度，使用Gemma 3 12B模型在WebVoyager和WebArena基准上达到开源最优水平。研究证明交互扩展是与计算扩展互补的重要维度，为训练自适应智能体开辟了新途径。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T09:02:26Z
- **目录日期**: 2025-06-10
