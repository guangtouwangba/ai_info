# Sim-and-Real Co-Training: A Simple Recipe for Vision-Based Robotic Manipulation

**URL**: http://arxiv.org/abs/2503.24361v1

## 原始摘要

Large real-world robot datasets hold great potential to train generalist
robot models, but scaling real-world human data collection is time-consuming
and resource-intensive. Simulation has great potential in supplementing
large-scale data, especially with recent advances in generative AI and
automated data generation tools that enable scalable creation of robot behavior
datasets. However, training a policy solely in simulation and transferring it
to the real world often demands substantial human effort to bridge the reality
gap. A compelling alternative is to co-train the policy on a mixture of
simulation and real-world datasets. Preliminary studies have recently shown
this strategy to substantially improve the performance of a policy over one
trained on a limited amount of real-world data. Nonetheless, the community
lacks a systematic understanding of sim-and-real co-training and what it takes
to reap the benefits of simulation data for real-robot learning. This work
presents a simple yet effective recipe for utilizing simulation data to solve
vision-based robotic manipulation tasks. We derive this recipe from
comprehensive experiments that validate the co-training strategy on various
simulation and real-world datasets. Using two domains--a robot arm and a
humanoid--across diverse tasks, we demonstrate that simulation data can enhance
real-world task performance by an average of 38%, even with notable differences
between the simulation and real-world data. Videos and additional results can
be found at https://co-training.github.io/


## AI 摘要

这篇论文探讨了如何结合仿真和真实世界数据来训练机器人策略。研究表明，仅靠仿真数据训练的策略迁移到现实世界需要大量人工调整，而将仿真数据与少量真实数据混合训练能显著提升性能。通过在不同机器人平台（机械臂和人形机器人）和任务上的实验，作者提出了一种简单有效的视觉操作任务训练方法。结果显示，即使仿真与现实数据存在明显差异，仿真数据仍能使现实任务表现平均提升38%。这项工作为利用仿真数据增强真实机器人学习提供了系统性见解和实践方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T01:28:34Z
- **目录日期**: 2025-04-02
