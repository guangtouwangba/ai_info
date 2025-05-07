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

该研究提出了一种新型的任务与运动规划（TAMP）方法，通过元优化技术连接高层规划与底层控制。该方法采用程序搜索优化轨迹作为基础模型与机器人控制的接口，并利用零阶方法优化模型输出的数值参数。相比现有方法（要么过度简化技能原语，要么直接预测关节角度），该方法在物体操控和绘图任务中表现更优，解决了抽象程度与精确控制的平衡问题。实验证明，该方法能有效提升规划效率并执行自然语言指令，为智能体与真实世界的交互提供了更优解决方案。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T23:01:26Z
- **目录日期**: 2025-05-07
