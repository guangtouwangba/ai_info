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

该研究探讨了结合仿真与真实世界数据共同训练机器人策略的有效方法。针对真实数据收集耗时耗资源的问题，研究者提出利用生成式AI和自动化工具创建大规模仿真数据集，并与少量真实数据混合训练。通过机器人手臂和人形机器人等多领域实验证明，这种协同训练策略能显著提升视觉操控任务的性能，平均提高38%，即使仿真与真实数据存在明显差异。研究提供了系统性的训练方法，验证了仿真数据对真实机器人学习的补充价值，为克服"现实差距"提供了可行方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-01T16:02:07Z
- **目录日期**: 2025-04-01
