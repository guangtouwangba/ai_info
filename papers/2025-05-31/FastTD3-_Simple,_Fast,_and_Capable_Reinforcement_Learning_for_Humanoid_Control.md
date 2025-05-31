# FastTD3: Simple, Fast, and Capable Reinforcement Learning for Humanoid Control

**URL**: http://arxiv.org/abs/2505.22642v2

## 原始摘要

Reinforcement learning (RL) has driven significant progress in robotics, but
its complexity and long training times remain major bottlenecks. In this
report, we introduce FastTD3, a simple, fast, and capable RL algorithm that
significantly speeds up training for humanoid robots in popular suites such as
HumanoidBench, IsaacLab, and MuJoCo Playground. Our recipe is remarkably
simple: we train an off-policy TD3 agent with several modifications -- parallel
simulation, large-batch updates, a distributional critic, and carefully tuned
hyperparameters. FastTD3 solves a range of HumanoidBench tasks in under 3 hours
on a single A100 GPU, while remaining stable during training. We also provide a
lightweight and easy-to-use implementation of FastTD3 to accelerate RL research
in robotics.


## AI 摘要

本文提出FastTD3算法，通过并行仿真、大批量更新、分布化评价器和精细调参等改进，显著加速人形机器人强化学习训练。该算法在HumanoidBench、IsaacLab等平台上，仅需单块A100 GPU不到3小时即可完成多种任务训练，且保持稳定性。研究提供了轻量级开源实现，为机器人强化学习研究提供高效工具。FastTD3以简洁的离策略TD3框架，解决了传统强化学习训练复杂耗时的痛点。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T11:02:06Z
- **目录日期**: 2025-05-31
