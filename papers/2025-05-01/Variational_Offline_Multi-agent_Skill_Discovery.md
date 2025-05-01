# Variational Offline Multi-agent Skill Discovery

**URL**: http://arxiv.org/abs/2405.16386v3

## 原始摘要

Skills are effective temporal abstractions established for sequential
decision making, which enable efficient hierarchical learning for long-horizon
tasks and facilitate multi-task learning through their transferability. Despite
extensive research, research gaps remain in multi-agent scenarios, particularly
for automatically extracting subgroup coordination patterns in a multi-agent
task. In this case, we propose two novel auto-encoder schemes: VO-MASD-3D and
VO-MASD-Hier, to simultaneously capture subgroup- and temporal-level
abstractions and form multi-agent skills, which firstly solves the
aforementioned challenge. An essential algorithm component of these schemes is
a dynamic grouping function that can automatically detect latent subgroups
based on agent interactions in a task. Further, our method can be applied to
offline multi-task data, and the discovered subgroup skills can be transferred
across relevant tasks without retraining. Empirical evaluations on StarCraft
tasks indicate that our approach significantly outperforms existing
hierarchical multi-agent reinforcement learning (MARL) methods. Moreover,
skills discovered using our method can effectively reduce the learning
difficulty in MARL scenarios with delayed and sparse reward signals. The
codebase is available at https://github.com/LucasCJYSDL/VOMASD.


## AI 摘要

该研究提出两种新型自动编码器方案（VO-MASD-3D和VO-MASD-Hier），首次实现多智能体系统中子群协调模式与时间抽象技能的自动提取。通过动态分组函数检测智能体交互中的潜在子群结构，该方法能从离线多任务数据中发现可跨任务迁移的子群技能。在《星际争霸》任务中的实验表明，该方法显著优于现有分层多智能体强化学习方法，并能有效缓解延迟稀疏奖励场景下的学习困难。研究成果为多智能体层级学习提供了新思路，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T11:02:12Z
- **目录日期**: 2025-05-01
