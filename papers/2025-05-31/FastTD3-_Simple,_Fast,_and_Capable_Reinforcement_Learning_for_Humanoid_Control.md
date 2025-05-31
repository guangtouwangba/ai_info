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

本文介绍了FastTD3算法，这是一种简单、快速且高效的强化学习方法，可显著加速人形机器人在HumanoidBench、IsaacLab和MuJoCo等平台的训练。该算法基于TD3框架，通过并行模拟、大批量更新、分布化critic网络及精细调参等改进，在单块A100 GPU上3小时内即可完成多种HumanoidBench任务，且训练过程稳定。研究团队还提供了轻量级、易用的FastTD3实现，以推动机器人强化学习研究的发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-31T13:08:35Z
- **目录日期**: 2025-05-31
