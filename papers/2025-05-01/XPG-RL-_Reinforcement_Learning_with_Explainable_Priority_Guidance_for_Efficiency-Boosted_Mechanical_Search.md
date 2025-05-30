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

XPG-RL是一种强化学习框架，用于解决杂乱环境中的机械搜索问题。它结合了可解释的优先级决策机制和情境感知切换策略，动态选择目标抓取、遮挡移除或视角调整等动作基元。其策略通过自适应阈值优化动作选择，感知模块融合RGB-D、语义和几何特征生成结构化场景表征。实验表明，XPG-RL在任务成功率和运动效率上显著优于基线方法，长时任务效率最高提升4.5倍，验证了结合领域知识与可学习策略对机器人操作的优越性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T02:32:24Z
- **目录日期**: 2025-05-01
