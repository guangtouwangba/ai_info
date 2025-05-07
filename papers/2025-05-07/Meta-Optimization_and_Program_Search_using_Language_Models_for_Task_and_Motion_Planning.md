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

该研究提出了一种新颖的任务与运动规划(TAMP)方法，通过元优化技术改进高层规划与低层控制的接口：(1)使用程序搜索结合轨迹优化作为基础模型与机器人控制的桥梁；(2)采用零阶方法优化基础模型输出的数值参数。相比现有方法(要么抽象过度依赖简化技能原语，要么缺乏抽象直接预测关节角度)，新方法在物体操纵和绘图等复杂任务中表现更优，同时保持了快速规划时间和自然语言指令执行能力。研究结果表明该方法优于先前的TAMP方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T18:01:37Z
- **目录日期**: 2025-05-07
