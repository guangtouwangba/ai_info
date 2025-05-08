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

本文提出了一种结合高层次任务规划和低层次运动控制的新型任务与运动规划(TAMP)方法。该方法通过程序搜索轨迹优化问题作为基础模型与机器人控制的接口，并采用零阶方法优化基础模型输出的数值参数，解决了现有方法在抽象程度上的局限性（要么过于简化技能原语，要么缺乏抽象）。实验表明，该方法在物体操作和绘图等复杂任务上优于现有TAMP方法，实现了更优的规划效果。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T01:29:06Z
- **目录日期**: 2025-05-08
