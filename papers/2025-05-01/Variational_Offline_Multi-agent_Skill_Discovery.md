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

该研究提出两种新型自动编码器方案（VO-MASD-3D和VO-MASD-Hier），首次实现多智能体任务中自动提取子群协调模式和时间抽象，形成可迁移的多智能体技能。核心创新是动态分组函数，能基于任务中的智能体交互自动检测潜在子群。该方法适用于离线多任务数据，发现的子群技能无需重新训练即可跨任务迁移。在《星际争霸》任务上的实验表明，其显著优于现有分层多智能体强化学习方法，并能有效缓解延迟和稀疏奖励场景下的学习困难。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T06:02:10Z
- **目录日期**: 2025-05-01
