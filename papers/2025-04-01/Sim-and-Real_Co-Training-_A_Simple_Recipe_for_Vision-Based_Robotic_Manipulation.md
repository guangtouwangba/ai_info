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

该研究探讨了如何结合仿真与真实世界数据共同训练机器人策略，以解决单纯仿真训练存在的"现实差距"问题。通过系统实验验证，研究者提出了一种简单有效的视觉机械臂操作任务训练方法。在机器人手臂和人形机器人两种平台上测试表明，即使仿真与真实数据存在显著差异，混合训练策略仍能平均提升38%的真实任务性能。这种方法突破了传统大规模真实数据收集的资源限制，为利用生成式AI和自动化工具创建可扩展的机器人行为数据集提供了新思路。研究结果展示了仿真数据对提升真实机器人学习效果的显著价值。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T20:01:14Z
- **目录日期**: 2025-04-01
