# XPG-RL: Reinforcement Learning with Explainable Priority Guidance for Efficiency-Boosted Mechanical Search

**URL**: http://arxiv.org/abs/2504.20969v1

## 原始摘要

Mechanical search (MS) in cluttered environments remains a significant
challenge for autonomous manipulators, requiring long-horizon planning and
robust state estimation under occlusions and partial observability. In this
work, we introduce XPG-RL, a reinforcement learning framework that enables
agents to efficiently perform MS tasks through explainable, priority-guided
decision-making based on raw sensory inputs. XPG-RL integrates a task-driven
action prioritization mechanism with a learned context-aware switching strategy
that dynamically selects from a discrete set of action primitives such as
target grasping, occlusion removal, and viewpoint adjustment. Within this
strategy, a policy is optimized to output adaptive threshold values that govern
the discrete selection among action primitives. The perception module fuses
RGB-D inputs with semantic and geometric features to produce a structured scene
representation for downstream decision-making. Extensive experiments in both
simulation and real-world settings demonstrate that XPG-RL consistently
outperforms baseline methods in task success rates and motion efficiency,
achieving up to 4.5$\times$ higher efficiency in long-horizon tasks. These
results underscore the benefits of integrating domain knowledge with learnable
decision-making policies for robust and efficient robotic manipulation.


## AI 摘要

本文提出XPG-RL强化学习框架，用于解决机械臂在杂乱环境中的搜索任务挑战。该框架结合任务驱动的动作优先级机制和情境感知切换策略，动态选择抓取、遮挡移除或视角调整等基础动作。其策略能自适应调整动作选择的阈值，并融合RGB-D输入与语义几何特征生成场景表征。实验表明，XPG-RL在任务成功率和运动效率上显著优于基线方法，长时任务效率提升高达4.5倍，验证了融合领域知识与可学习决策策略对提升机器人操作鲁棒性和效率的有效性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T22:01:37Z
- **目录日期**: 2025-04-30
