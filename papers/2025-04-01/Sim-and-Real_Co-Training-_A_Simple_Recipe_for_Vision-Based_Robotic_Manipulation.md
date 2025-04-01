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

这篇论文探讨了如何结合仿真和真实世界数据来训练机器人策略。研究表明，尽管大规模真实数据收集成本高昂，但通过混合使用仿真数据和少量真实数据进行协同训练，能显著提升视觉机械臂任务的性能。实验证明，即使仿真与真实数据存在明显差异，这种方法仍能使任务表现平均提升38%。研究提供了有效的协同训练方法，并在机械臂和人形机器人等多种任务上验证了其效果。该方法为利用仿真数据增强真实机器人学习提供了实用方案，相关成果详见项目网站。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T15:01:20Z
- **目录日期**: 2025-04-01
