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

该研究提出了一种新颖的任务与运动规划（TAMP）方法，通过结合高层规划与底层控制来优化机器人操作。该方法采用元优化技术，利用程序搜索轨迹优化问题作为基础模型与机器人控制的接口，并通过零阶方法优化模型输出的数值参数。相比现有方法（如简化技能链或直接预测关节角度），该方法在物体操控和绘图等复杂任务中表现更优，解决了抽象过度或不足的问题。实验证明，该方法显著提升了TAMP性能，同时支持自然语言指令执行和快速规划。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T00:01:46Z
- **目录日期**: 2025-05-08
