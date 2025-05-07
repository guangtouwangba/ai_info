# Meta-Optimization and Program Search using Language Models for Task and Motion Planning

**URL**: http://arxiv.org/abs/2505.03725v1

## 原始摘要

Intelligent interaction with the real world requires robotic agents to
jointly reason over high-level plans and low-level controls. Task and motion
planning (TAMP) addresses this by combining symbolic planning and continuous
trajectory generation. Recently, foundation model approaches to TAMP have
presented impressive results, including fast planning times and the execution
of natural language instructions. Yet, the optimal interface between high-level
planning and low-level motion generation remains an open question: prior
approaches are limited by either too much abstraction (e.g., chaining
simplified skill primitives) or a lack thereof (e.g., direct joint angle
prediction). Our method introduces a novel technique employing a form of
meta-optimization to address these issues by: (i) using program search over
trajectory optimization problems as an interface between a foundation model and
robot control, and (ii) leveraging a zero-order method to optimize numerical
parameters in the foundation model output. Results on challenging object
manipulation and drawing tasks confirm that our proposed method improves over
prior TAMP approaches.


## AI 摘要

该研究提出了一种新颖的任务与运动规划(TAMP)方法，通过元优化技术改进高层规划与底层控制的接口。该方法：(1)使用程序搜索结合轨迹优化作为基础模型与机器人控制的桥梁；(2)采用零阶方法优化基础模型输出的数值参数。相比现有方法(要么过于抽象如简单技能链，要么缺乏抽象如直接预测关节角度)，新方法在物体操控和绘图等复杂任务上表现更优。实验证实该方法提升了TAMP系统的性能，平衡了抽象程度与控制精度，为解决机器人智能交互中的规划与控制联合推理问题提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T21:01:23Z
- **目录日期**: 2025-05-07
