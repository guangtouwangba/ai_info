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

XPG-RL是一种强化学习框架，用于解决杂乱环境中的机械搜索（MS）问题。它结合了可解释的优先级决策机制和上下文感知切换策略，动态选择目标抓取、遮挡移除和视角调整等动作基元。该框架通过优化策略输出自适应阈值来指导动作选择，并利用RGB-D输入与语义几何特征融合的场景表示进行决策。实验表明，XPG-RL在任务成功率和运动效率上显著优于基线方法，长时任务效率提升高达4.5倍，验证了结合领域知识与可学习策略在机器人操作中的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T23:01:39Z
- **目录日期**: 2025-04-30
