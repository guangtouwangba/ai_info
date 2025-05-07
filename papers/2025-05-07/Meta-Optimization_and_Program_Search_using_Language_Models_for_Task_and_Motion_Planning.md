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

该研究提出了一种新型任务与运动规划(TAMP)方法，通过元优化技术解决高层规划与底层控制间的接口问题。方法采用程序搜索轨迹优化作为基础模型与机器人控制的接口，并利用零阶方法优化模型输出的数值参数。相比现有方法(过于抽象或缺乏抽象)，该方法在物体操纵和绘图任务中表现更优，结合了符号规划与连续轨迹生成的优点，同时支持自然语言指令。实验结果验证了其优于传统TAMP方法的性能。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T05:01:30Z
- **目录日期**: 2025-05-07
